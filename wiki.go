package main

import (
	"bufio"
	"io/ioutil"
	"regexp"
	"strings"
)

type link struct {
	uid  string
	name string
	url  string
}

// parseSummaryFile parses GitBook Summary.md file and returns links.
func parseSummaryFile() []link {
	bytes, _ := ioutil.ReadFile("summary.md")

	// regex to extract markdown links
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)
	re1 := regexp.MustCompile(`(.md)`)

	var links []link

	// Read file line by line and extraxt links
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		links = append(links, link{
			uid:  matches[0][1],
			name: matches[0][1],
			url:  "https://wiki.nikitavoloboev.xyz/" + re1.ReplaceAllString(matches[0][2], `.html`),
		})
	}
	return links
}
