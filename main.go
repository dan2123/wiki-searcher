package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	defaultLimit = 5
	historyLimit = 5
)

func printResults(wr WikiResponse) {
	fmt.Printf("Query Term:\t%v\n", wr.QueryTerm)
	fmt.Printf("Timestamp:\t%v\n", time.Now().Format(time.RFC3339))
	for i := range wr.ResultTitles {
		fmt.Printf("Title: %v, Desc: %v, Link: %v\n", wr.ResultTitles[i], wr.ResultDescriptions[i], wr.ResultLinks[i])
	}
}

func handleError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	// initialize cache to store history
	cache, err := NewLRUCache(historyLimit)
	if err != nil {
		handleError(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(">>> ")
	searchTerm, err := reader.ReadString('\n')
	if err != nil {
		handleError(err)
	}

	for {
		searchTerm = strings.TrimSuffix(searchTerm, "\n")
		cache.Add(searchTerm)
		fmt.Println("Search History:\t", cache.GetItems())

		res, err := SearchWiki(searchTerm, defaultLimit)
		if err != nil {
			handleError(err)
		}
		printResults(res)

		fmt.Printf(">>> ")
		searchTerm, err = reader.ReadString('\n')
		if err != nil {
			handleError(err)
		}
	}
}
