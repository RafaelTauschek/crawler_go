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
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	cfg.mu.Lock()
	pagesLen := len(cfg.pages)
	cfg.mu.Unlock()

	if pagesLen >= cfg.maxPages {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if cfg.baseUrl.Hostname() != currentURL.Hostname() {
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
		baseUrl := cfg.baseUrl.String()

		urls, err := getURLsFromHTML(html, baseUrl)
		if err != nil {
			return
		}

		for _, url := range urls {
			fmt.Printf("Started to crawl URL: %v\n", url)
			cfg.wg.Add(1)
			go cfg.crawlPage(url)
		}
	}
}
