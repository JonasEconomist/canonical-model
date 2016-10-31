package content

var ArticleSample = Content{
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
	URLCollection: URLCollection{
		Canonical: "http://www.economist.com/news/21697812-worlds-most-valuable-company-reported-its-first-year-year-quarterly-revenue-decline",
		Web:       "/news/21697812-worlds-most-valuable-company-reported-its-first-year-year-quarterly-revenue-decline",
		Short:     "bit.ly/21697812",
	},
	Headline:    "Shake it off",
	SubHeadline: "The future of Apple",
	Description: "The worldâ€™s most valuable company needs another mega hit",
	PrintEdition: Print{
		SubHeadline: "The worldâ€™s most valuable company",
		Edition:     "http://mt-content.stage.s.aws.economist.com/mapper/id/21697795",
		Section:     "http://mt-content.stage.s.aws.economist.com/sections/77",
		PageStart:   25,
		PageEnd:     26,
	},
	IsPartOf: []ListLinks{
		ListLinks{
			Position: "2",
			List:     "abcjrusqjuiub84fu6t3h3n6oivvaa1b",
		},
	},
	ArticleSection: Section{
		Internal: []Ref{
			NewRef(ID("/sections/34")),
		},
	},
	About: []AboutLink{
		AboutLink{
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
	Ads: Ads{
		Zone:      "kjdu",
		Site:      "LASN",
		Grapeshot: "http://mt-content.stage.s.aws.economist.com/external/grapeshot?url=http%3A%2F%2Fwww.economist.com%2Fnode%2F21697812",
	},
	Sponsor: Person{
		GivenName:  "Mark",
		FamilyName: "Brincat",
	},
	Comment: NewRef(ID("/node/21697812/comments#comments")),
	Images: Images{
		Main: "http://cdn.static-economist.com/sites/default/files/images/2016/04/articles/body/20160430_wbc249_0.png",
		Inline: []WebURL{
			"http://cdn.static-economist.com/sites/default/files/images/2016/04/articles/body/20160430_wbc249_0.png",
		},
	},
	Author: []Person{
		Person{
			FamilyName: "M.S.R",
		},
	},
	Text: Body{
		MediaType: "text/html",
		Text:      "<p>OUR product pipeline has amazing innovations in store.</p>",
	},
	Publications: []Publication{
		Publication{
			SubHeadline: "The worldâ€™s most valuable company",
		},
	},
}

var HomepageSample = Content{
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
	URLCollection: URLCollection{
		Canonical: "http://www.economist.com/",
		Web:       "/",
	},
	Headline: "The Economist - World News, Politics, Economics, Business & Finance",
	HasPart:  "http://localhost:9494/list/abcjrusqjuiub84fu6t3h3n6oivvaa1b",
	Genre: []string{
		"News",
		"Business",
		"Economics",
	},
	Ads: Ads{
		Zone:      "kjdu",
		Site:      "LASN",
		Grapeshot: "http://mt-content.stage.s.aws.economist.com/external/grapeshot?url=http%3A%2F%2Fwww.economist.com%2Fnode%2F21697812",
	},
	Sponsor: Person{
		GivenName:  "Mark",
		FamilyName: "Brincat",
	},
}

var CollectionSample = Content{
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
	URLCollection: URLCollection{
		Canonical: "http://www.economist.com/node/21701829",
		Web:       "/node/21701829",
	},
	Headline: "Story collection 2 - new",
	HasPart:  "http://localhost:9494/list/efgjrusqjuiub84fu6t3h3n6oivvaa1b",
	IsPartOf: []ListLinks{
		ListLinks{
			Position: "1",
			List:     "abcjrusqjuiub84fu6t3h3n6oivvaa1b",
		},
	},
	About: []AboutLink{
		AboutLink{
			Source: "Topics",
			Ref:    "http://mt-content.stage.s.aws.economist.com/lists/topics/21701829",
		},
	},
	Genre: []string{
		"News",
		"Business",
		"Economics",
	},
	Ads: Ads{
		Zone:      "kjdu",
		Site:      "LASN",
		Grapeshot: "http://mt-content.stage.s.aws.economist.com/external/grapeshot?url=http%3A%2F%2Fwww.economist.com%2Fnode%2F21701829",
	},
	Sponsor: Person{
		GivenName:  "Mark",
		FamilyName: "Brincat",
	},
	Publications: []Publication{
		Publication{
			SubHeadline: "Included!",
		},
	},
}
