package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joshakeman/morning_post"
)

func main() {
	links, err := morning_post.HNLinks(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	for _, link := range links {
		fmt.Println(link)
	}
}

