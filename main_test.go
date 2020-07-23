package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestGetURL(t *testing.T) {
	t.Parallel()
	want := "https://www.works-hub.com/jobs/remote-senior-go-engineer-e69?utm_source=Linkedin&utm_medium=Recruiter_Social&utm_campaign=p.gubbey"
	got, err := GetURL("&amp;#32; submitted by &amp;#32; &lt;a href=&quot;https://www.reddit.com/user/Pmgubbey1&quot;&gt; /u/Pmgubbey1 &lt;/a&gt; &lt;br/&gt; &lt;span&gt;&lt;a href=&quot;https://www.works-hub.com/jobs/remote-senior-go-engineer-e69?utm_source=Linkedin&amp;amp;utm_medium=Recruiter_Social&amp;amp;utm_campaign=p.gubbey&quot;&gt;[link]&lt;/a&gt;&lt;/span&gt; &amp;#32; &lt;span&gt;&lt;a href=&quot;https://www.reddit.com/r/golang/comments/hil0ox/senior_remote_golang_job/&quot;&gt;[comments]&lt;/a&gt;&lt;/span&gt;")
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

// NOTES:
// Got test passing, since entrylist has content it is passing
