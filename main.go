package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
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

func (c Client) getFeed(url string) (Feed, error) {
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return Feed{}, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Feed{}, err
	}
	var feed Feed
	if err = xml.Unmarshal(data, &feed); err != nil {
		return Feed{}, fmt.Errorf("decoding xml %q: %v", data, err)
	}
	return feed, nil
}

type Feed struct {
	XMLName   xml.Name `xml:"feed"`
	EntryList []Entry  `xml:entry`
}

type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Title   string   `xml:"title"`
	Content string   `xml:"content"`
}
