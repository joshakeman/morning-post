package main

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
)

func main() {
	body, err := getFeed("https://www.reddit.com/r/golang/.rss?format=xml")
	if err != nil {
		log.Println(err)
	}

	marshalXML(body)
}

func getFeed(url string) (io.ReadCloser, error) {
	resp, err := http.Get("https://www.reddit.com/r/golang/.rss?format=xml")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	return resp.Body, nil
}

func marshalXML(body io.ReadCloser) {
	var feed Feed

	z, err := xml.Marshal(body)
	if err != nil {
		log.Println(err)
	}

	xml.Unmarshal(z, &feed)
}

type Feed struct {
	XMLName   xml.Name `xml:"feed"`
	EntryList []Entry  `xml:"entry"`
}

type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Title   string   `xml:"title"`
	Content string   `xml:"content,innerxml"`
}
