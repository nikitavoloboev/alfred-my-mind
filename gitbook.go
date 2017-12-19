// Package gitbook contains helper functions to parse GitBook summaries
package main

import (
	"bufio"
	"io/ioutil"
	"regexp"
	"strings"
)

// parseSummary parses GitBook Summary.md file
// TODO: add special case for Introduction (one entry)
func parseSummary() {
	bytes, _ := ioutil.ReadFile("Summary.md")

	// regex to extract markdown links
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)
	re1 := regexp.MustCompile(`(.md)`)

	// Read string line by line and apply regex
	// TODO: add error check
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		wf.NewItem(matches[0][1]).Arg("https://wiki.nikitavoloboev.xyz/" + re1.ReplaceAllString(matches[0][2], `.html`)).Valid(true).UID(matches[0][1])
	}
}
