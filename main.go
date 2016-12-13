package main

import (
	"context"
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.google.co.uk",
	"http://golang.org",
	"http://www.bbc.co.uk",
	"https://uk.yahoo.com/",
}

func req(ctx context.Context, url string, back chan string) {
	tr := &http.Transport{}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp.Body.Close()
	back <- url
}

func makeRequests() {
	ctx, cancel := context.WithCancel(context.Background())
	back := make(chan string)
	fmt.Println("Making requests...")
	for _, url := range urls {
		go req(ctx, url, back)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
		case u := <-back:
			fmt.Printf("%s wins!\n", u)
			cancel()
			return
		}
	}
}

func main() {
	makeRequests()
}
