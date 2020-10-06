package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

func main() {
	c := NewClient()
	reader, err := c.getFeed("https://hnrss.org/newest?q=Golang")
	if err != nil {
		log.Println(err)
	}
	links, err := ReadFeedFrom(reader)
	if err != nil {
		log.Println(err)
	}
	for _, v := range links.Channel.Items {
		log.Printf(v.Link)
	}
}

type Client struct {
	HTTPClient *http.Client
}

func NewClient() Client {
	return Client{
		HTTPClient: http.DefaultClient,
	}
}

func (c Client) getFeed(url string) (io.Reader, error) {
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// feed, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return resp.Body{}, err
	// }
	// var feed HNfeed
	// if err = xml.Unmarshal(data, &feed); err != nil {
	// 	return HNfeed{}, fmt.Errorf("decoding xml %q: %v", data, err)
	// }
	return resp.Body, nil
}

func ReadFeedFrom(r io.Reader) (HNfeed, error) {
	var feed HNfeed
	if err := xml.NewDecoder(r).Decode(&feed); err != nil {
		return HNfeed{}, err
	}
	return feed, nil
}

func GetURLs(f HNfeed) ([]string, error) {
	return []string{}, nil
}

/* TYPES */

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
