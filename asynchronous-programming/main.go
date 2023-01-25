package main

import (
	"fmt"
	"net/http"
)

type urlStatus struct {
	url    string
	status bool
}

func checkUrl(url string, c chan urlStatus) {
	_, err := http.Get(url)
	if err != nil {
		// The website is down
		c <- urlStatus{url, false}
	} else {
		c <- urlStatus{url, true}
	}
}

func main() {
	urls := []string{
		"https://www.easyjet.com/",
		"https://www.skyscanner.de/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
		"https://www.swiss.com/",
	}

	c := make(chan urlStatus)
	for _, url := range urls {
		go checkUrl(url, c)
	}

	result := make([]urlStatus, len(urls))
	for i := range result {
		result[i] = <-c
		if result[i].status {
			fmt.Println(result[i].url, "is up")
		} else {
			fmt.Println(result[i].url, "is down")
		}
	}
}
