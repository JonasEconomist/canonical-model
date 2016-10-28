package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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
	bytes, err := json.Marshal(articleSample)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(bytes)))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(homepageSample)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(string(bytes)))
}

func collection(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(collectionSample)
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

func main() {
	http.HandleFunc("/", api)
	http.HandleFunc("/article/", article)
	http.HandleFunc("/homepage/", homepage)
	http.HandleFunc("/collection/", collection)
	http.HandleFunc("/populate/", populate)
	http.ListenAndServe(":9494", nil)
}

var articleSample = content.Content{
	ID:                   "http://mt-content.stage.s.aws.economist.com/mapper/id/21697812",
	TegID:                "thpjrusqjuiub84fu6t3h3n6oivvaa1b",
	TegType:              "article",
	Type:                 "article",
	Version:              "0",
	RegionsAllowed:       []string{"GB"},
	RequiresSubscription: false,
	IsAccessibleForFree:  true,
	DateCreated:          "2015-12-30T15:49:20Z",
	DateModified:         "2015-12-30T15:49:20Z",
	DatePublished:        "2015-12-30T15:49:20Z",
	InLanguage:           "en",
	URLCollection: content.URLCollection{
		Canonical: "http://www.economist.com/news/21697812-worlds-most-valuable-company-reported-its-first-year-year-quarterly-revenue-decline",
		Web:       "/news/21697812-worlds-most-valuable-company-reported-its-first-year-year-quarterly-revenue-decline",
		Short:     "bit.ly/21697812",
	},
	Headline:    "Shake it off",
	SubHeadline: "The future of Apple",
	Description: "The worldâ€™s most valuable company needs another mega hit",
	PrintEdition: content.Print{
		SubHeadline: "The worldâ€™s most valuable company",
		Edition:     "http://mt-content.stage.s.aws.economist.com/mapper/id/21697795",
		Section:     "http://mt-content.stage.s.aws.economist.com/sections/77",
		PageStart:   25,
		PageEnd:     26,
	},
	IsPartOf: []content.ListLinks{
		content.ListLinks{
			Position: "2",
			List:     "abcjrusqjuiub84fu6t3h3n6oivvaa1b",
		},
	},
	ArticleSection: content.Section{
		Internal: []content.Ref{
			content.NewRef(content.ID("/sections/34")),
		},
	},
	About: []content.AboutLink{
		content.AboutLink{
			Source: "Topics",
			Ref:    "http://mt-content.stage.s.aws.economist.com/lists/topics/21697816",
		},
	},
	Genre: []string{
		"News",
		"Business",
		"Economics",
	},
	Keywords: []string{
		"Business",
		"Apple",
	},
	Ads: content.Ads{
		Zone:      "kjdu",
		Site:      "LASN",
		Grapeshot: "http://mt-content.stage.s.aws.economist.com/external/grapeshot?url=http%3A%2F%2Fwww.economist.com%2Fnode%2F21697812",
	},
	Sponsor: content.Person{
		GivenName:  "Mark",
		FamilyName: "Brincat",
	},
	Comment: content.NewRef(content.ID("/node/21697812/comments#comments")),
	Images: content.Images{
		Main: "http://cdn.static-economist.com/sites/default/files/images/2016/04/articles/body/20160430_wbc249_0.png",
		Inline: []content.WebURL{
			"http://cdn.static-economist.com/sites/default/files/images/2016/04/articles/body/20160430_wbc249_0.png",
		},
	},
	Author: []content.Person{
		content.Person{
			FamilyName: "M.S.R",
		},
	},
	Text: content.Body{
		MediaType: "text/html",
		Text:      "<p>OUR product pipeline has amazing innovations in store.</p>",
	},
	Publications: []content.Publication{
		content.Publication{
			SubHeadline: "The worldâ€™s most valuable company",
		},
	},
}
var req, _ = http.NewRequest("POST", "http://localhost:9200/article/testmap/_search", bytes.NewBufferString(`{ "query": { "match": { "isPartOf.list": "abcjrusqjuiub84fu6t3h3n6oivvaa1b" } }}`))

var homepageSample = content.Content{
	ID:                   "http://mt-content.stage.s.aws.economist.com/mapper/id/21555491",
	TegID:                "abcjrusqjuiub84fu6t3h3n6oivvaa1b",
	TegType:              "homepage",
	Type:                 "homepage",
	Version:              "0",
	RegionsAllowed:       []string{"GB"},
	RequiresSubscription: false,
	IsAccessibleForFree:  true,
	DateCreated:          "2015-12-30T15:49:20Z",
	DateModified:         "2015-12-30T15:49:20Z",
	DatePublished:        "2015-12-30T15:49:20Z",
	InLanguage:           "en",
	URLCollection: content.URLCollection{
		Canonical: "http://www.economist.com/",
		Web:       "/",
	},
	Headline: "The Economist - World News, Politics, Economics, Business & Finance",
	HasPart:  req,
	Genre: []string{
		"News",
		"Business",
		"Economics",
	},
	Ads: content.Ads{
		Zone:      "kjdu",
		Site:      "LASN",
		Grapeshot: "http://mt-content.stage.s.aws.economist.com/external/grapeshot?url=http%3A%2F%2Fwww.economist.com%2Fnode%2F21697812",
	},
	Sponsor: content.Person{
		GivenName:  "Mark",
		FamilyName: "Brincat",
	},
}

var req2, _ = http.NewRequest("POST", "http://localhost:9200/article/testmap/_search&q=isPartOf:efgjrusqjuiub84fu6t3h3n6oivvaa1b", nil)

var collectionSample = content.Content{
	ID:                   "http://mt-content.stage.s.aws.economist.com/collections/21701829",
	TegID:                "efgjrusqjuiub84fu6t3h3n6oivvaa1b",
	TegType:              "storyCollection",
	Type:                 "storyCollection",
	Version:              "0",
	RegionsAllowed:       []string{"GB"},
	RequiresSubscription: false,
	IsAccessibleForFree:  true,
	DateCreated:          "2015-12-30T15:49:20Z",
	DateModified:         "2015-12-30T15:49:20Z",
	DatePublished:        "2015-12-30T15:49:20Z",
	InLanguage:           "en",
	URLCollection: content.URLCollection{
		Canonical: "http://www.economist.com/node/21701829",
		Web:       "/node/21701829",
	},
	Headline: "Story collection 2 - new",
	HasPart:  req2,
	IsPartOf: []content.ListLinks{
		content.ListLinks{
			Position: "1",
			List:     "abcjrusqjuiub84fu6t3h3n6oivvaa1b",
		},
	},
	About: []content.AboutLink{
		content.AboutLink{
			Source: "Topics",
			Ref:    "http://mt-content.stage.s.aws.economist.com/lists/topics/21701829",
		},
	},
	Genre: []string{
		"News",
		"Business",
		"Economics",
	},
	Ads: content.Ads{
		Zone:      "kjdu",
		Site:      "LASN",
		Grapeshot: "http://mt-content.stage.s.aws.economist.com/external/grapeshot?url=http%3A%2F%2Fwww.economist.com%2Fnode%2F21701829",
	},
	Sponsor: content.Person{
		GivenName:  "Mark",
		FamilyName: "Brincat",
	},
	Publications: []content.Publication{
		content.Publication{
			SubHeadline: "Included!",
		},
	},
}
