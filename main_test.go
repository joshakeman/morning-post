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
	if err != nil {
		t.Fatal(err)
	}
	if len(feed.Channel.Items) == 0 {
		t.Fatal("EntryList is empty")
	}
}

// func TestGetURL(t *testing.T) {
// 	t.Parallel()

// 	tcs := []struct {
// 		input string
// 		want  string
// 	}{
// 		{
// 			input: "&amp;#32; submitted by &amp;#32; &lt;a href=&quot;https://www.reddit.com/user/Pmgubbey1&quot;&gt; /u/Pmgubbey1 &lt;/a&gt; &lt;br/&gt; &lt;span&gt;&lt;a href=&quot;https://www.works-hub.com/jobs/remote-senior-go-engineer-e69?utm_source=Linkedin&amp;amp;utm_medium=Recruiter_Social&amp;amp;utm_campaign=p.gubbey&quot;&gt;[link]&lt;/a&gt;&lt;/span&gt; &amp;#32; &lt;span&gt;&lt;a href=&quot;https://www.reddit.com/r/golang/comments/hil0ox/senior_remote_golang_job/&quot;&gt;[comments]&lt;/a&gt;&lt;/span&gt;",
// 			want:  "https://www.works-hub.com/jobs/remote-senior-go-engineer-e69?utm_source=Linkedin&utm_medium=Recruiter_Social&utm_campaign=p.gubbey",
// 		},
// 		{
// 			input: "&lt;!-- SC_OFF --&gt;&lt;div class=&quot;md&quot;&gt;&lt;p&gt;I am writing a multi-coroutine task processing engine. Anyone want with me together.&lt;/p&gt; &lt;p&gt;I have 2+ years experience with golang, but i am just like a beginner. Maybe this is the philosophy of golang.&lt;/p&gt; &lt;p&gt;&lt;a href=&quot;https://github.com/90634/gotaskengine&quot;&gt;https://github.com/90634/gotaskengine&lt;/a&gt;&lt;/p&gt; &lt;p&gt;This is my first post. Is this correct behavior?&lt;/p&gt; &lt;/div&gt;&lt;!-- SC_ON --&gt; &amp;#32; submitted by &amp;#32; &lt;a href=&quot;https://www.reddit.com/user/dafsic&quot;&gt; /u/dafsic &lt;/a&gt; &lt;br/&gt; &lt;span&gt;&lt;a href=&quot;https://www.reddit.com/r/golang/comments/hifopt/a_sample_task_engine/&quot;&gt;[link]&lt;/a&gt;&lt;/span&gt; &amp;#32; &lt;span&gt;&lt;a href=&quot;https://www.reddit.com/r/golang/comments/hifopt/a_sample_task_engine/&quot;&gt;[comments]&lt;/a&gt;&lt;/span&gt;",
// 			want:  "https://github.com/90634/gotaskengine",
// 		},
// 	}

// 	for _, tc := range tcs {
// 		got, err := GetURL(tc.input)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		if !cmp.Equal(tc.want, got) {
// 			t.Error(cmp.Diff(tc.want, got))
// 		}
// 	}
// }

// NOTES:
// Got test passing, since entrylist has content it is passing
