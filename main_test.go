package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
	f, err := os.Open("testdata/hn.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	got, err := ReadFeedFrom(f)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
