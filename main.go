package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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

func canonical(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(sample)
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

func readData(path string, f os.FileInfo, err error) error {
	log.Fatalf("% #v\n", f)
	return nil
}

func main() {
	http.HandleFunc("/", api)
	http.HandleFunc("/content/", canonical)
	http.HandleFunc("/populate/", populate)
	http.ListenAndServe(":9494", nil)
}

var sample = content.Content{
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
	IsPartOf: []content.Ref{
		content.NewRef(content.ID("/lists/print")),
		content.NewRef(content.ID("/lists/sections")),
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
	Comment: content.NewRef(content.ID("http://www.economist.com/node/21697812/comments#comments")),
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
	ListData: content.List{
		WithinLists: []content.Ref{
			content.NewRef(content.ID("/mapper/id/21697795")),
			content.NewRef(content.ID("/sections/77")),
			content.NewRef(content.ID("/sections/34")),
		},
	},
}
