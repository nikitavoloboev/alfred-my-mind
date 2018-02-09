// Package gitbook contains helper functions to parse GitBook summaries.
package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func doSearch() error {
	showUpdateStatus()

	log.Printf("query=%s", query)

	parseSummary()

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()

	return nil
}

// parseSummary parses GitBook Summary.md file
// TODO: Add special case for Introduction (one entry)
func parseSummary() {
	bytes, _ := ioutil.ReadFile("Summary.md")

	// regex to extract markdown links
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)
	re1 := regexp.MustCompile(`(.md)`)

	// Read string line by line and apply regex
	// TODO: add map error check
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		wf.NewItem(matches[0][1]).Arg("https://wiki.nikitavoloboev.xyz/" + re1.ReplaceAllString(matches[0][2], `.html`)).Valid(true).UID(matches[0][1])
	}
}
