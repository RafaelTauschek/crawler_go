package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseUrl            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %v\n", args[0])

	url, err := url.Parse(args[0])
	if err != nil {
		fmt.Println("can't parse url")
		os.Exit(1)
	}

	cfg := config{
		pages:   make(map[string]int),
		baseUrl: url,
	}

	cfg.crawlPage(args[0])

	for url, count := range cfg.pages {
		fmt.Printf("%s: %d\n", url, count)
	}

}
