package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
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
	time_start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp.Body.Close()
	back <- fmt.Sprintf("%s - %s", time.Since(time_start).String(), url)
}

func makeRequests() {
	//ctx, _ := context.WithCancel(context.Background())
	//ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)
	ctx := context.Background()
	back := make(chan string)
	fmt.Println("Making requests...")
	for _, url := range urls {
		go req(ctx, url, back)
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		case u := <-back:
			fmt.Println(u)
		//cancel()
		//return
		}
	}
}

func main() {
	makeRequests()
}
