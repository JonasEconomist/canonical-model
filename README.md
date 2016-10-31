# Economist Canonical Data Model

Run the main.go file to have to API documentation and a sample content served on your localhost port 9494 at article/, homepage/, and collection/.

This code provides full RAML documentation and a rough outline of the Go structures that encapsulate the Economist canonical data model as outlined in the spreadsheet found here: 
https://docs.google.com/a/economist.com/spreadsheets/d/1eyFVfnu6pKlW58QGXc7cevE3o9MvNrTO40MGremJLfw/edit?usp=sharing

Additional samples can be added to illustration how other types of content, such as images, lists, and blog sites, would fit into this model.

# Canonical Data Model Elastic Search Mappings

This codebase also contains a script to create an index and mappings
for an Elastic Search cluster, either running locally or on our [stage cluster](mt-content-search.s.aws.economist.com).

* To install ES on OSX run ```brew install elasticsearch```.
* Run ```elasticsearch``` to boot the node.
* To set up the ```article``` index and populate with some sample data, run ```./setup.sh```. Be warned, this will tear down the indices and rebuild them.
* To populate the index with sample data, fill up the ```data``` directory with as many documents as you like, fire up the service and send a request to the ```populate``` endpoint. You need to have Elastic Search running for this to work, and the data must align to the canonical data model. Any errors will be logged to the console and the service will return a 500. 

## Search

Here are some sample queries you can run after initial setup to get you going:

* view mappings for the index: ```http://localhost:9200/article/_mapping```
* search for the term "clinton": ```http://localhost:9200/ article/_search?q=clinton```

* search for the term "clinton" in ```testmap``` types: ````http://localhost:9200/article/testmap/_search?q=clinton```

* search for the term "clinton" and return the headline and canonical URL: ```http://localhost:9200/article/testmap/_search?q=clinton&fields=headline,url.canonical```

* filtered search (which doesn't order by score and so is faster and easier to cache): ```http://localhost:9200/article/_search``` make a POST request with a body of:

```json
{
	"query": {
		"filtered": {
			"filter": {
				"term": {
					"name": "trump"
				}
			}
		}
	}
}
```

## Content Relationships

The lists/{id} endpoint runs a query against the elasticsearch service to find all content that is a part of the supplied content tegID. To illustration how this works, navigate to the homepage/ endpoint and click on the link in the hasPart field.

This will return the results of an elastichsearch query to find all content that has the submitted tegID in the isPartOf field. The tegIDs of this content is returned by the search. View the article/ and collection/ endpoint to verify that both have the homepage tegID in the isPartOf field.

