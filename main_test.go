package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetFeed(t *testing.T) {
	t.Parallel()
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		f, err := os.Open("testdata/hn.xml")
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		io.Copy(w, f)
	}))
	c := NewClient()
	c.HTTPClient = ts.Client()
	feed, err := c.getFeed(ts.URL)
	log.Println(feed)
	if err != nil {
		t.Fatal(err)
	}
	// if len(feed.Channel.Items) == 0 {
	// 	t.Fatal("EntryList is empty")
	// }
}

func TestReadFeedFrom(t *testing.T) {
	t.Parallel()

	want := HNfeed{
		Channel: Channel{
			Items: []Item{
				{Link: "https://www.reddit.com/r/golang/comments/56oirg/api_encoding_shootout_binary_data_protobuf_vs/"},
				{Link: "https://github.com/ansd/lastpass-go"},
				{Link: "https://www.pixelstech.net/article/1599275392-Implement-struct-no-copy-in-GoLang"},
			},
		},
	}

	r := strings.NewReader(`<rss version="2.0" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:atom="http://www.w3.org/2005/Atom">
							<channel>
								<item><link>https://www.reddit.com/r/golang/comments/56oirg/api_encoding_shootout_binary_data_protobuf_vs/</link></item>
								<item><link>https://github.com/ansd/lastpass-go</link></item>
								<item><link>https://www.pixelstech.net/article/1599275392-Implement-struct-no-copy-in-GoLang</link></item>
							</channel>
							</rss>`)
	got, err := ReadFeedFrom(r)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

// func TestGetURL(t *testing.T) {
// 	t.Parallel()

// 	// tcs := []struct {
// 	// 	input HNfeed
// 	// 	want  string
// 	// }{
// 	// 	input
// 	// }

// 	want := []string{
// 		"https://www.reddit.com/r/golang/comments/56oirg/api_encoding_shootout_binary_data_protobuf_vs/",
// 	}
// 	feed := getFeed()
// 	got, err := GetURLs()
// 	// for _, tc := range tcs {
// 	// 	got, err := GetURLs(tc.input)
// 	// 	if err != nil {
// 	// 		t.Fatal(err)
// 	// 	}
// 	// 	if !cmp.Equal(tc.want, got) {
// 	// 		t.Error(cmp.Diff(tc.want, got))
// 	// 	}
// 	// }
// }

// NOTES:
// Got test passing, since entrylist has content it is passing
