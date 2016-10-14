package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

func handler(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.Marshal(sample)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(string(bytes)))
}

func main() {
	http.HandleFunc("/", api)
	http.HandleFunc("/content/", handler)
	http.ListenAndServe(":9494", nil)
}

// Content represents a canonical Content post.
type Content struct {
	ID                   Ref            `json:"id"`
	TegID                string         `json:"tegID"`
	TegType              string         `json:"tegType"`
	Type                 string         `json:"type"`
	SameAs               string         `json:"sameAs,omitempty"`
	Version              string         `json:"version"`
	RegionsAllowed       []string       `json:"regionsAllowed,omitempty"`
	RequiresSubscription bool           `json:"requiresSubscription,omitempty"`
	IsAccessibleForFree  bool           `json:"isAccessibleForFree,omitempty"`
	DateCreated          ISODate        `json:"dateCreated"`
	DateModified         ISODate        `json:"dateModified,omitempty"`
	DatePublished        ISODate        `json:"datePublished,omitempty"`
	DateExpired          ISODate        `json:"dateExpired,omitempty"`
	InLanguage           string         `json:"inLanguage"`
	URLCollection        URLCollection  `json:"url,omitempty"`
	Headline             string         `json:"headline"`
	SubHeadline          string         `json:"subheadline,omitempty"`
	Description          string         `json:"description,omitempty"`
	Byline               string         `json:"byline,omitempty"`
	Tagline              string         `json:"tagline,omitempty"`
	PrintEdition         Print          `json:"printEdition,omitempty"`
	Publisher            Person         `json:"publisher,omitempty"`
	Channel              string         `json:"channel,omitempty"`
	IsPartOf             []Ref          `json:"isPartOf,omitempty"`
	ArticleSection       Section        `json:"articleSection,omitempty"`
	About                []AboutLink    `json:"about,omitempty"`
	Genre                []string       `json:"genre,omitempty"`
	Keywords             []string       `json:"keywords,omitempty"`
	Ads                  Ads            `json:"ads,omitempty"`
	Sponsor              Person         `json:"sponsor,omitempty"`
	Audience             []Person       `json:"audience,omitempty"`
	Comment              Ref            `json:"comment,omitempty"`
	Images               Images         `json:"images,omitempty"`
	Videos               Videos         `json:"videos,omitempty"`
	Audio                Audio          `json:"audio,omitempty"`
	AssociatedMedia      []Ref          `json:"associatedMedia,omitempty"`
	MediaData            Media          `json:"mediaData,omitempty"`
	Author               []Person       `json:"author,omitempty"`
	Contributor          []Person       `json:"contributor,omitempty"`
	CopyrightHolder      []Person       `json:"copyrightHolder,omitempty"`
	AccountablePerson    []Person       `json:"accountablePerson,omitempty"`
	LocationCreated      []Location     `json:"locationCreated,omitempty"`
	SourceOrganization   []Organization `json:"sourceOrganization,omitempty"`
	CopyrightYear        int            `json:"copyrightYear,omitempty"`
	License              string         `json:"license,omitempty"`
	Text                 Body           `json:"text,omitempty"`
	Publications         []Publication  `json:"publications"`
	ListData             List           `json:"list,omitempty"`
}

// Publication holds the overridden values for content.
type Publication struct {
	Headline    string `json:"headline,omitempty"`
	SubHeadline string `json:"subHeadline,omitempty"`
	Description string `json:"description,omitempty"`
}

// List represents a re-usable list of hypermedia links.
type List struct {
	Name        string     `json:"name,omitempty"`
	Position    int        `json:"position,omitempty"`
	Order       string     `json:"order,omitempty"`
	TotalItems  int        `json:"totalItems,omitempty"`
	Pagination  Pagination `json:"pagination,omitempty"`
	List        Ref        `json:"list,omitempty"`
	WithinLists []Ref      `json:"withinLists,omitempty"`
}

// URLCollection holds a collection of URLs for a piece of content.
type URLCollection struct {
	Canonical WebURL `json:"canonical,omitempty"`
	Web       WebURL `json:"web,omitempty"`
	Short     string `json:"short,omitempty"`
	Twitter   WebURL `json:"twitter,omitempty"`
	Alias     WebURL `json:"alias,omitempty"`
}

// NewURLCollection is factory function for URLCollection type.
func NewURLCollection(rawURL string, shortURL string, finder PathFinder) (urlColl URLCollection) {
	urlColl.Web = NewWebURL(rawURL, finder)
	urlColl.Short = shortURL

	return urlColl
}

// Print holds a collection of data relavent to articles in the Print Edition.
type Print struct {
	Headline    string `json:"headline,omitempty"`
	SubHeadline string `json:"subHeadline,omitempty"`
	Description string `json:"description,omitempty"`
	Byline      string `json:"byline,omitempty"`
	Tagline     string `json:"tagline,omitempty"`
	Section     Ref    `json:"section,omitempty"`
	Edition     Ref    `json:"edition,omitempty"`
	PageStart   int    `json:"pageStart,omitempty"`
	PageEnd     int    `json:"pageEnd,omitempty"`
	Column      string `json:"column,omitempty"`
}

type Person struct {
	GivenName  string `json:"givenName,omitempty"`
	FamilyName string `json:"familyName,omitempty"`
}

// Section holds a collection of internal and external section data.
type Section struct {
	Internal []Ref    `json:"internal,omitempty"`
	External []string `json:"external,omitempty"`
}

// Ads holds a collection of relevant ad data.
type Ads struct {
	Zone      string `json:"zone,omitempty"`
	Site      string `json:"site,omitempty"`
	Grapeshot Ref    `json:"grapeshot,omitempty"`
}

// Images holds a collection Images associated with the content.
type Images struct {
	Main    WebURL   `json:"main,omitempty"`
	Inline  []WebURL `json:"internal,omitempty"`
	Gallery []WebURL `json:"gallery,omitempty"`
}

// Videos holds a collection Videos associated with the content.
type Videos struct {
	Main    WebURL   `json:"main,omitempty"`
	Inline  []WebURL `json:"internal,omitempty"`
	Gallery []WebURL `json:"gallery,omitempty"`
}

// Audio holds a collection Images associated with the content.
type Audio struct {
	Main   WebURL   `json:"main,omitempty"`
	Inline []WebURL `json:"internal,omitempty"`
}

// Media contains additional media information for supported media types.
type Media struct {
	Width          int    `json:"width,omitempty"`
	Height         int    `json:"height,omitempty"`
	AutoStart      bool   `json:"autostart,omitempty"`
	EncodingFormat string `json:"encodingFormat,omitempty"`
	Bitrate        int    `json:"bitrate,omitempty"`
	Rendition      []Ref  `json:"rendition,omitempty"`
}

// Ads holds the format and text of the content body.
type Body struct {
	MediaType string `json:"mediaType,omitempty"`
	Text      string `json:"text,omitempty"`
}

type Organization struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
}

type Location struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
}

// Pagination represents field in struct with information about range of elements in list.
type Pagination struct {
	PageStart int `json:"pageStart,omitempty"`
	PageEnd   int `json:"pageEnd,omitempty"`
}

// AboutLink holds hypermedia link with a source description.
type AboutLink struct {
	Source string `json:"source,omitempty"`
	Ref    Ref    `json:"ref,omitempty"`
}

// NewAboutLink is the factory function for AboutLink structure.
func NewAboutLink(source string, ref Ref) AboutLink {
	return AboutLink{Source: source, Ref: ref}
}

// Ref represents a consistently output Hypermedia Link.
type Ref string

// NewRef is the factory function for the Ref type.
func NewRef(i Identifier) (ref Ref) {
	id := i.GetID()
	if _, err := strconv.Atoi(id); err == nil {
		ref = Ref(fmt.Sprintf("%s%s%s", os.Getenv("API_GATEWAY_ENDPOINT"), "/mapper/id/", id))
	} else {
		ref = Ref(fmt.Sprintf("%s%s", os.Getenv("API_GATEWAY_ENDPOINT"), id))
	}
	return ref
}

// Identifier represents how to identify a piece of content or list.
type Identifier interface {
	GetID() string
}

// ID represents a content or list ID and is used to create hypermedia.
// links for that ID.
type ID string

// GetID - implementation of Identifier interface.
func (i ID) GetID() string {
	return string(i)
}

// WebURL represents a consistently output URL.
type WebURL string

// NewWebURL is the factory function for the WebURL type.
func NewWebURL(rawURL string, finder PathFinder) WebURL {
	switch true {
	case rawURL == "":
		return WebURL(finder.GetPath())
	case strings.HasPrefix(rawURL, "/"):
		return WebURL(rawURL)
	default:
		return WebURL("/" + rawURL)
	}
}

// PathFinder is an interface used for values to
// tell services about their original locations if a web
// alias cannot be found i.e. /node/<nid> for Drupal.
type PathFinder interface {
	GetPath() string
}

// ISODate represents ISO 8601 (RFC 3339) dates.
// See https://en.wikipedia.org/wiki/ISO_8601.
type ISODate string

// String is the ISODate Stringer implementation.
func (i ISODate) String() string {
	return string(i)
}

// NewISODate is the factory function for ISODates. It can be passed
// a timestamp int or a string, in the latter case it will try and
// convert the string into an integer before doing the conversion.
func NewISODate(datetime interface{}) (isoDate ISODate, err error) {
	var data int64
	switch datetime.(type) {
	case int:
		data = int64(datetime.(int))
	case string:
		data, err = strconv.ParseInt(datetime.(string), 10, 32)
		if err != nil {
			return isoDate, errors.New("unable to convert date string")
		}
	}
	iso := ISODate(time.Unix(data, 0).UTC().Format(time.RFC3339))
	if iso.String() == "1970-01-01T00:00:12Z" {
		return isoDate, errors.New("unable to format date string")
	}
	return iso, nil
}

var sample = Content{
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
	IsPartOf: []Ref{
		NewRef(ID("http://mt-content.stage.s.aws.economist.com/lists/print")),
		NewRef(ID("http://mt-content.stage.s.aws.economist.com/lists/sections")),
	},
	ArticleSection: Section{
		Internal: []Ref{
			NewRef(ID("http://mt-content.stage.s.aws.economist.com/sections/34")),
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
	Comment: NewRef(ID("http://www.economist.com/node/21697812/comments#comments")),
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
	ListData: List{
		WithinLists: []Ref{
			NewRef(ID("http://mt-content.stage.s.aws.economist.com/mapper/id/21697795")),
			NewRef(ID("http://mt-content.stage.s.aws.economist.com/sections/77")),
			NewRef(ID("http://mt-content.stage.s.aws.economist.com/sections/34")),
		},
	},
}
