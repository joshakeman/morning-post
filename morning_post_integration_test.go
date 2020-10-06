//+build integration

package morning_post_test

import (
	"strings"
	"testing"

	"github.com/joshakeman/morning_post"
)

func TestHNLinks(t *testing.T) {
	got, err := morning_post.HNLinks("golang")
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Fatal("want some links, got none")
	}
	if !strings.HasPrefix(got[0], "http") {
		t.Fatalf("want 'http' at start, got %q", got[0])
	}
}