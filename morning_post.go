package morning_post

import (
	"encoding/xml"
	"io"
	"net/http"
)

type HNfeed struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Link string `xml:"link"`
}

func ReadFeedFrom(r io.Reader) (HNfeed, error) {
	var feed HNfeed
	if err := xml.NewDecoder(r).Decode(&feed); err != nil {
		return HNfeed{}, err
	}
	return feed, nil
}

func GetURLs(f HNfeed) (URLs []string) {
	for _, item := range f.Channel.Items {
		URLs = append(URLs, item.Link)
	}
	return URLs
}

func HNLinks(query string) ([]string, error) {
	resp, err := http.Get("https://hnrss.org/newest?q="+query)
	if err != nil {
		return []string{}, err
	}
	defer resp.Body.Close()
	feed, err := ReadFeedFrom(resp.Body)
	if err != nil {
		return []string{}, err
	}
	return GetURLs(feed), nil
}
