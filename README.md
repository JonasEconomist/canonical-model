# Economist Canonical Data Model

Run the main.go file to have to API documentation and a sample article served on your localhost port 9494 at content/.

This code provides full RAML documentation and a rough outline of the Go structures that encapsulate the Economist canonical data model as outlined in the spreadsheet found here: 
https://docs.google.com/a/economist.com/spreadsheets/d/1eyFVfnu6pKlW58QGXc7cevE3o9MvNrTO40MGremJLfw/edit?usp=sharing

Additional samples can be added to illustration how other types of content, such as images, lists, and blog sites, would fit into this model.

# Canonical Data Model Elastic Search Mappings

This codebase also contains some a script to create an index and mappings
for an Elastic Search cluster, either running locally or on our [stage cluster](mt-content-search.s.aws.economist.com).

To set up the ```article``` index and populate with some sample data, run
```./populate.sh```. Be warned, this will tear down the indices and rebuild them.

To populate the index with sample data, fill up the ```data``` directory with as many documents as you like, fire up the service and send a request
to the ```populate``` endpoint. You need to have Elastic Search running for this to work, and the data must align to the canonical data model. Any errors will be logged to the console and the service will return a 500. 
