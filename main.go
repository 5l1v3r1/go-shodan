package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shodan-api/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Enter a search term")
	}

	// Please do not hardcode API keys!!!
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.ProfileInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf(
		"Query Credits: %d\nScan Credits:  %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	profileSearch, err := s.ProfileSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	for _, host := range profileSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
	// Hits the exploit api, this will be the first argument passed to it
	xPloit, err := s.ExploitSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	for _, exploit := range xPloit.Matches {
		fmt.Printf("Author: %s\n\nVuln Info: %s\n\nSource: %s\n\nCVE: %s",
			exploit.Author,
			exploit.Description,
			exploit.Source,
			exploit.Cve)
	}
}
