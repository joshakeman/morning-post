package main

import (
	"io"
	"log"
	"testing"
)

func TestGetFeed(t *testing.T) {
	body, err := getFeed("https://www.reddit.com/r/golang/.rss?format=xml")
	if err != nil {
		log.Println(err)
	}

	_, ok := body.(io.ReadCloser)
	if ok != true {
		t.Errorf("getFeed did not return an io.ReadCloser")
	}
}
