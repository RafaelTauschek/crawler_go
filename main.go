package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	maxPages           int
	baseUrl            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	url, err := url.Parse(args[0])
	if err != nil {
		os.Exit(1)
	}

	concurrency, err := strconv.Atoi(args[1])
	if err != nil {
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		os.Exit(1)
	}

	cfg := &config{
		pages:              make(map[string]int),
		maxPages:           maxPages,
		baseUrl:            url,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, concurrency),
		wg:                 &sync.WaitGroup{},
	}

	fmt.Printf("starting crawl of: %v\n", args[0])

	cfg.wg.Add(1)
	go cfg.crawlPage(args[0])
	cfg.wg.Wait()

	printReport(cfg.pages, cfg.baseUrl.String())
}
