package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseUrl string) {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %v\n", baseUrl)
	fmt.Println("=============================")

	sortedPages := sortPages(pages)

	for _, page := range sortedPages {
		fmt.Printf("Found %d internal links to %v\n", page.count, page.url)
	}
}

type Page struct {
	url   string
	count int
}

func sortPages(pages map[string]int) []Page {

	pageSlice := make([]Page, 0, len(pages))

	for k, val := range pages {
		pageSlice = append(pageSlice, Page{url: k, count: val})
	}

	sort.Slice(pageSlice, func(i, j int) bool {
		if pageSlice[i].count != pageSlice[j].count {
			return pageSlice[i].count > pageSlice[j].count
		}

		return pageSlice[i].url < pageSlice[j].url
	})

	return pageSlice
}
