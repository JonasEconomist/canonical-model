package es

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/JonasEconomist/canonical-model/content"
	"github.com/olivere/elastic"
)

func ReadData() error {
	var content content.Content
	d, err := os.Open("./data")
	if err != nil {
		return err
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		return err
	}

	fmt.Println("Reading data from data directory")

	client, err := elastic.NewClient()
	if err != nil {
		return err
	}

	fmt.Println("Elasticsearch client booted")

	for _, file := range files {
		if file.Mode().IsRegular() {
			data, err := ioutil.ReadFile("./data/" + file.Name())
			if err != nil {
				log.Printf("file reading error: %v\n", err)
				return err
			}
			fmt.Printf("Reading data from %s\n", file.Name())
			if err := json.Unmarshal(data, &content); err != nil {
				log.Printf("JSON decoding error: %v\n", err)
				return err
			}
			fmt.Printf("Decoded JSON into content value\n")

			_, err = client.Index().
				Index("article").
				Type("testmap").
				Id(content.TegID).
				BodyJson(content).
				Refresh(true).
				Do()
			if err != nil {
				log.Printf("Indexing error: %v\n", err)

				return err
			}
			fmt.Println("content successfully indexed")
		}
	}
	return nil
}
