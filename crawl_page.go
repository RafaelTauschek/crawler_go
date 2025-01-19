package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	_, ok := cfg.pages[normalizedURL]

	if ok {
		cfg.mu.Lock()
		cfg.pages[normalizedURL]++
		cfg.mu.Unlock()
		return false
	}

	cfg.mu.Lock()
	cfg.pages[normalizedURL] = 1
	cfg.mu.Unlock()
	return true
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if cfg.baseUrl.Host != currentURL.Host {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	isFirst := cfg.addPageVisit(normalizedURL)

	if isFirst {
		html, err := getHTML(rawCurrentURL)
		if err != nil {
			return
		}

		fmt.Printf("Recieved HTML from current URL: %v\n", normalizedURL)

		baseUrl := cfg.baseUrl.String()

		urls, err := getURLsFromHTML(html, baseUrl)
		if err != nil {
			return
		}

		fmt.Printf("Found %d URLS on %s\n", len(urls), normalizedURL)
		fmt.Printf("URLs found: %v\n", urls)

		for _, url := range urls {
			fmt.Printf("Started to crawl URL: %v\n", url)
			cfg.crawlPage(url)
		}
	}
}
