package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
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

func ReadFeedFrom(r io.Reader) (HNfeed, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return HNfeed{}, err
	}
	var feed HNfeed
	if err = xml.Unmarshal(data, &feed); err != nil {
		return HNfeed{}, fmt.Errorf("decoding xml %q: %v", data, err)
	}
	return feed, nil
}

func GetURLs(f HNfeed) ([]string, error) {
	return []string{}, nil
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
	// XMLName xml.Name `xml:"rss"`
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Link string `xml:"link"`
}
