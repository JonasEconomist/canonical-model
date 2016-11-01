package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/JonasEconomist/canonical-model/content"
	"github.com/JonasEconomist/canonical-model/es"
)

func api(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("api.html")
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(data))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func article(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(content.ArticleSample)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(bytes)))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(content.HomepageSample)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(bytes)))
}

func collection(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(content.CollectionSample)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(bytes)))
}

func populate(w http.ResponseWriter, r *http.Request) {
	err := es.ReadData()
	if err != nil {
		log.Printf("error reading data: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func list(w http.ResponseWriter, r *http.Request) {
	idRegex := regexp.MustCompile("^.*?/list/(.*)$")
	id := idRegex.FindStringSubmatch(r.URL.Path)
	query := fmt.Sprintf(`{ "sort": [{"isPartOf.position": {"order": "asc"}}], "_source" : ["tegID", "tegType", "headline", "isPartOf.position"], "query": { "match": { "isPartOf.list": "%v" } }}`, id[1])
	req, err := http.NewRequest("POST", "http://localhost:9200/article/testmap/_search", strings.NewReader(query))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(body)))
}

func main() {
	http.HandleFunc("/", api)
	http.HandleFunc("/article/", article)
	http.HandleFunc("/homepage/", homepage)
	http.HandleFunc("/collection/", collection)
	http.HandleFunc("/list/", list)
	http.HandleFunc("/populate/", populate)
	http.ListenAndServe(":9494", nil)
}
