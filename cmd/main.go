package main

import (
	"fmt"

	"github.com/joshakeman/morning_post"
)

func main() {
	links := morning_post.HNLinks("golang")
	fmt.Println(links)
}

