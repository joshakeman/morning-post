package main

import (
	"encoding/xml"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.reddit.com/r/golang/.rss?format=xml")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	content := xml.NewDecoder(resp.Body)
	log.Println(content)
}
