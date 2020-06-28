package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// resp, err := http.Get("https://www.reddit.com/r/golang/.rss?format=xml")
	// if err != nil {
	// 	log.Println(err)
	// }
	// defer resp.Body.Close()

	// content := xml.NewDecoder(resp.Body)
	// log.Println(content)
	body, err := getFeed("https://www.reddit.com/r/golang/.rss?format=xml")
	if err != nil {
		log.Println(err)
	}
	log.Println(body)
}

func getFeed(url string) (io.ReadCloser, error) {
	resp, err := http.Get("https://www.reddit.com/r/golang/.rss?format=xml")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	return resp.Body, nil
}
