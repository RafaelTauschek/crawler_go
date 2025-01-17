package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if baseURL.Host != currentURL.Host {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	_, ok := pages[normalizedURL]
	if ok {
		pages[normalizedURL]++
	} else {
		pages[normalizedURL] = 1

		html, err := getHTML(rawCurrentURL)
		if err != nil {
			return
		}
		fmt.Printf("Recieved HTML from current URL: %v\n", normalizedURL)

		urls, err := getURLsFromHTML(html, rawBaseURL)
		if err != nil {
			return
		}

		fmt.Printf("Found %d URLS on %s\n", len(urls), normalizedURL)
		fmt.Printf("URLs found: %v\n", urls)

		for _, url := range urls {
			fmt.Printf("Started to crawl URL: %v\n", url)
			crawlPage(rawBaseURL, url, pages)
		}
	}
}
