package main

import (
	"encoding/xml"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/antchfx/htmlquery"
)

// func main() {
// 	_, err := getFeed("https://www.reddit.com/r/golang/.rss?format=xml")
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

type Client struct {
	HTTPClient *http.Client
}

func NewClient() Client {
	return Client{
		HTTPClient: http.DefaultClient,
	}
}

func (c Client) getFeed(url string) (HNfeed, error) {
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return HNfeed{}, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	log.Println(len(data))
	if err != nil {
		return HNfeed{}, err
	}
	var feed HNfeed
	if err = xml.Unmarshal(data, &feed); err != nil {
		return HNfeed{}, fmt.Errorf("decoding xml %q: %v", data, err)
	}
	return feed, nil
}

func GetURL(s string) (string, error) {
	esc := html.UnescapeString(s)

	doc, err := htmlquery.Parse(strings.NewReader(esc))
	if err != nil {
		return "", err
	}
	a := htmlquery.Find(doc, "//a@href")

	// for _, v := range a {
	// 	log.Println(*v)
	// }
	// a2 := htmlquery.FindOne(doc, "//a[1]@href")
	href := htmlquery.SelectAttr(a[1], "href")
	return href, nil
}

type Link struct {
	Href string `xml:"href,attr"`
}

type Feed struct {
	XMLName   xml.Name `xml:"feed"`
	EntryList []Entry  `xml:"entry"`
}

type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Title   string   `xml:"title"`
	Content string   `xml:"content"`
	Link    Link     `xml:"link"`
}

// Created new Link struct that's embedded in Entry, allowing us to pull href value from that link.

type HNfeed struct {
	XMLName xml.Name `xml:"rss"`
	// Text    string   `xml:",chardata"`
	// Version string   `xml:"version,attr"`
	// Dc      string   `xml:"dc,attr"`
	// Atom    string   `xml:"atom,attr"`
	Channel struct {
		// Text  string `xml:",chardata"`
		// Title string `xml:"title"`
		// Link  struct {
		// 	Text string `xml:",chardata"`
		// 	Href string `xml:"href,attr"`
		// 	Rel  string `xml:"rel,attr"`
		// 	Type string `xml:"type,attr"`
		// } `xml:"link"`
		// Description   string `xml:"description"`
		// Docs          string `xml:"docs"`
		// Generator     string `xml:"generator"`
		// LastBuildDate string `xml:"lastBuildDate"`
		Items []struct {
			// Text        string `xml:",chardata"`
			// Title       string `xml:"title"`
			// Description string `xml:"description"`
			// PubDate     string `xml:"pubDate"`
			Link string `xml:"link"`
			// Creator     string `xml:"creator"`
			// Comments    string `xml:"comments"`
			// Guid        struct {
			// 	Text        string `xml:",chardata"`
			// 	IsPermaLink string `xml:"isPermaLink,attr"`
			// } `xml:"guid"`
		} `xml:"item"`
	} `xml:"channel"`
}
