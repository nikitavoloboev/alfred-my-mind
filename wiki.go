package main

import (
	"bufio"
	"io/ioutil"
	"regexp"
	"strings"
)

// Link holds a link with unique UID.
type Link struct {
	uid  string
	name string
	url  string
}

// parseSummaryFile parses GitBook Summary.md file and returns links.
func parseSummaryFile() []Link {
	bytes, _ := ioutil.ReadFile("SUMMARY.md")

	// regex to extract markdown links
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)
	// re1 := regexp.MustCompile(`(.md)`)

	var links []Link

	// Read file line by line and extraxt links
	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		split := strings.Split(matches[0][2], "/")

		noMD := strings.Split(split[len(split)-1], ".")[0]
		// Delete last item if it's equal to the pre last item in paths
		if split[len(split)-2] == noMD {
			split = split[:len(split)-1]
		}
		// Remove .md if present
		last := split[len(split)-1]
		if strings.Contains(last, ".md") {
			last = last[:len(last)-3]
		}
		split[len(split)-1] = last
		if contains(split, "macOS") {
			for i, v := range split {
				split[0] = "macos"
				if v == "apps" {
					split[i] = "macos-apps"
				}
			}
		}

		links = append(links, Link{
			uid:  matches[0][1],
			name: matches[0][1],
			url:  "https://wiki.nikitavoloboev.xyz/" + strings.Join(split, "/"),
		})
	}
	return links
}

// check if string is included in a slice
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func searchWiki() {
	showUpdateStatus()
	links := parseSummaryFile()

	// TODO: implement caching
	for _, link := range links {
		wf.NewItem(link.name).UID(link.uid).Valid(true).Arg(link.url)
	}

	// TODO: message doesnt show
	wf.WarnEmpty("No matching items", "Try a different query?")

	wf.SendFeedback()
}
