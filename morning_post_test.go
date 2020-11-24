package morning_post_test

import (
	"os"
	"testing"

	"github.com/joshakeman/morning_post"

	"github.com/google/go-cmp/cmp"
)

func TestReadFeedFrom(t *testing.T) {
	t.Parallel()

	want := morning_post.HNfeed{
		Channel: morning_post.Channel{
			Items: []morning_post.Item{
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
	got, err := morning_post.ReadFeedFrom(f)
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetURLs(t *testing.T) {
	input := morning_post.HNfeed{
		Channel: morning_post.Channel{
			Items: []morning_post.Item{
				{Link: "https://www.reddit.com/r/golang/comments/56oirg/api_encoding_shootout_binary_data_protobuf_vs/"},
				{Link: "https://github.com/ansd/lastpass-go"},
				{Link: "https://www.pixelstech.net/article/1599275392-Implement-struct-no-copy-in-GoLang"},
			},
		},
	}
	want := []string{
		"https://www.reddit.com/r/golang/comments/56oirg/api_encoding_shootout_binary_data_protobuf_vs/",
		"https://github.com/ansd/lastpass-go",
		"https://www.pixelstech.net/article/1599275392-Implement-struct-no-copy-in-GoLang",
	}
	got := morning_post.GetURLs(input)
	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}

