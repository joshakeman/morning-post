package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetFeed(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		f, err := os.Open("testdata/feed.xml")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		io.Copy(w, f)
	}))
	c := NewClient()
	c.HTTPClient = ts.Client()
	feed, err := c.getFeed(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if len(feed.EntryList) == 0 {
		t.Fatal("EntryList is empty")
	}
}
