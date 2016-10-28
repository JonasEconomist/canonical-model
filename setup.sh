#!/usr/bin/env bash

ADDRESS=$1

if [ -z $ADDRESS ]; then
  ADDRESS="localhost:9200"
fi

# Check that Elasticsearch is running.
curl -s "http://$ADDRESS" 2>&1 > /dev/null
if [ $? != 0 ]; then
    echo "Unable to contact Elasticsearch at $ADDRESS"
    echo "Please ensure Elasticsearch is running and can be reached at http://$ADDRESS/"
    exit -1
fi

echo "WARNING, this script will delete the 'article' index and re-index all data!"
echo "Press Control-C to cancel this operation."
echo
echo "Press [Enter] to continue."
read

# Delete the old index.
curl -s -XDELETE "$ADDRESS/article" > /dev/null

# Create the next index from mapping file.
echo "Creating 'article' index..."
curl -s -XPOST "$ADDRESS/article" -d@$(dirname $0)/elasticsearchmap.json
curl -s "$ADDRESS/article/_health?wait_for_status=yellow&timeout=10s" > /dev/null
echo
echo "Done creating 'article' index."

echo
echo "Indexing data..."

echo "Indexing homepage..."
curl -s -XPOST "$ADDRESS/article/testmap/abcjrusqjuiub84fu6t3h3n6oivvaa1b" -d'{
  "id": "http://mt-content.stage.s.aws.economist.com/mapper/id/21555491",
  "tegID": "abcjrusqjuiub84fu6t3h3n6oivvaa1b",
  "tegType": "homepage",
  "type": "homepage",
  "version": "0",
  "regionsAllowed": ["GB"],
  "isAccessibleForFree": true,
  "dateCreated": "2016-10-23T15:49:20Z",
  "dateModified": "2016-10-23T15:49:20Z",
  "datePublished": "2016-10-23T15:49:20Z",
  "inLanguage": "en",
  "url": {
    "canonical": "http://www.economist.com/",
    "web": "/",
    "short": ""
  },
  "headline": "The Economist - World News, Politics, Economics, Business & Finance",
  "subheadline": "",
  "description": "",
  "hasPart": "http://localhost:9200/article/testmap/_search&q=isPartOf:abcjrusqjuiub84fu6t3h3n6oivvaa1b",
  "isPartOf": "",
  "articleSection": {},
  "about": [],
  "genre": ["News", "Politics", "Economics", "Business", "Finance"],
  "keywords": [],
  "ads": {
    "zone": "kjdu",
    "site": "LASN",
    "grapeshot": "http://mt-content.stage.s.aws.economist.com/external/grapeshot?url=http%3A%2F%2Fwww.economist.com"
  },
  "sponsor": {
    "givenName": "Mark",
    "familyName": "Brincat"
  },
  "comment": "",
  "images": {},
  "author": [{
    "familyName": "M.S.R"
  }],
  "text": {},
  "publications": [],
  "list": {
    "withinLists": []
  }
}'

echo
echo "Indexing article..."
curl -s -XPOST "$ADDRESS/article/testmap/thpjrusqjuiub84fu6t3h3n6oivvaa1b" -d'{
  "id": "http://mt-content.stage.s.aws.economist.com/mapper/id/21707839",
  "tegID": "thpjrusqjuiub84fu6t3h3n6oivvaa1b",
  "tegType": "blog",
  "type": "article",
  "version": "0",
  "requiresSubscription": true,
  "dateCreated": "2016-09-26T13:22:31Z",
  "dateModified": "2016-09-26T19:03:23Z",
  "datePublished": "2016-09-26T15:16:26Z",
  "inLanguage": "en",
  "url": {
    "canonical": "http://www.economist.com/blogs/democracyinamerica/2016/09/clinton-v-trump",
    "web": "//blogs/democracyinamerica/2016/09/clinton-v-trump",
    "short": "bit.ly/21707839"
  },
  "headline": "What difference do presidential debates make?",
  "subheadline": "Clinton v Trump",
  "description": "Debates tend not to have a huge effect on the polls, but this is an unusual year",
  "hasPart": "",
  "isPartOf": ["abcjrusqjuiub84fu6t3h3n6oivvaa1b"],
  "articleSection": {
    "internal": ["http://mt-content.stage.s.aws.economist.comhttp://mt-content.stage.s.aws.economist.com/sections/77911"]
  },
  "about": [{
    "source": "Topics",
    "ref": "http://mt-content.stage.s.aws.economist.com/lists/topics/21707839"
  }],
  "genre": ["News", "Politics", "United States"],
  "keywords": ["Trump", "Clinton"],
  "ads": {
    "zone": "none",
    "site": "FMSQ",
    "grapeshot": "http://mt-content.stage.s.aws.economist.com/external/grapeshot?url=http%3A%2F%2Fwww.economist.com%2Fnode%2F21707839"
  },
  "sponsor": {
    "givenName": "Mark",
    "familyName": "Brincat"
  },
  "comment": "http://mt-content.stage.s.aws.economist.comhttp://www.economist.com/node/21707839/comments#comments",
  "images": {
    "main": "http://cdn.static-economist.com/sites/default/files/20161001_USP501_473.jpg",
    "internal": ["http://cdn.static-economist.com/sites/default/files/20160924_WOC988_0.png"]
  },
  "author": [{
    "familyName": "W.Z."
  }],
  "text": {
    "mediaType": "text/html",
    "text": "\u003cp\u003eHILLARY CLINTON and Donald Trump are the most unpopular presidential candidates ever.\u003c/p\u003e"
  },
  "publications": [{
    "subHeadline": "This is an unusual year"
  }],
  "list": {
    "pagination": {},
    "withinLists": ["http://mt-content.stage.s.aws.economist.com/lists/blogs/21003976", "http://mt-content.stage.s.aws.economist.com/sections/77911", "http://mt-content.stage.s.aws.economist.com/sections/77952"]
  }
}'

echo
echo "Indexing story collection..."
curl -s -XPOST "$ADDRESS/article/testmap/efgjrusqjuiub84fu6t3h3n6oivvaa1b" -d'{
  "id": "http://mt-content.stage.s.aws.economist.com/collections/21701829",
  "tegID": "efgjrusqjuiub84fu6t3h3n6oivvaa1b",
  "tegType": "storyCollection",
  "type": "list",
  "version": "0",
  "dateCreated": "2015-12-30T15:49:20Z",
  "dateModified": "2015-12-30T15:49:20Z",
  "datePublished": "2015-12-30T15:49:20Z",
  "inLanguage": "en",
  "url": {
    "canonical": "http://www.economist.com/node/21701829",
    "web": "/node/21701829"
  },
  "headline": "Story collection 2 - new",
  "hasPart": "http://localhost:9200/article/testmap/_search&q=isPartOf:efgjrusqjuiub84fu6t3h3n6oivvaa1b",
  "isPartOf": ["abcjrusqjuiub84fu6t3h3n6oivvaa1b"],
  "about": [{
    "source": "Topics",
    "ref": "http://mt-content.stage.s.aws.economist.com/lists/topics/21701829"
  }],
  "genre": ["United States", "Elections", "Politics"],
  "keywords": ["Politics", "Clinton"],
  "ads": {
    "zone": "kjdu",
    "site": "LASN",
    "grapeshot": "http://mt-content.stage.s.aws.economist.com/external/grapeshot?url=http%3A%2F%2Fwww.economist.com%2Fnode%2F21701829"
  },
  "list": {
    "order": "descending",
    "pagination": {
      "pageStart": 0,
      "pageEnd": 3
    },
    "list": ["http://mt-content.stage.s.aws.economist.com/mapper/id/21697812", "http://mt-content.stage.s.aws.economist.com/mapper/id/21708420", "http://mt-content.stage.s.aws.economist.com/mapper/id/21708069"],
    "withinLists": ["http://mt-content.stage.s.aws.economist.comhttp://mt-content.stage.s.aws.economist.com/mapper/id/21697795", "http://mt-content.stage.s.aws.economist.com", "http://mt-content.stage.s.aws.economist.com/sections/34"]
  }
}'

echo
echo "Indexing section..."
curl -s -XPOST "$ADDRESS/article/testmap/hijjrusqjuiub84fu6t3h3n6oivvaa1b" -d'{
  "id": "http://mt-content.stage.s.aws.economist.com/sections/34",
  "tegID": "hijjrusqjuiub84fu6t3h3n6oivvaa1b",
  "tegType": "section",
  "type": "section",
  "version": "0",
  "dateCreated": "2015-12-30T15:49:20Z",
  "dateModified": "2015-12-30T15:49:20Z",
  "datePublished": "2015-12-30T15:49:20Z",
  "inLanguage": "en",
  "url": {
    "canonical": "http://www.economist.com/sections/business-finance",
    "web": "sections/sections/business-finance"
  },
  "headline": "Business",
  "hasPart": "http://localhost:9200/article/testmap/_search&q=isPartOf:hijjrusqjuiub84fu6t3h3n6oivvaa1b",
  "isPartOf": ["http://mt-content.stage.s.aws.economist.com/lists/sections"],
  "genre": ["News", "Business", "Finance"],
  "keywords": ["Business"],
  "ads": {
    "zone": "kjdu",
    "site": "LASN"
  },
  "list": {
    "withinLists": ["http://mt-content.stage.s.aws.economist.com/lists/sections"]
  }
}'

echo
echo "Done indexing content."

# Refresh so data is available
curl -s -XPOST "$ADDRESS/article/_refresh"

echo
echo "Done indexing data."
echo
