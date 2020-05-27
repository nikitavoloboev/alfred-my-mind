package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

type bookmarks map[string]string

func searchLinks(wikiPath string) {
	showUpdateStatus()
	bookmarks := make(bookmarks)

	files := parseSummary()
	for _, file := range files {
		parseLinks(wikiPath, file, bookmarks)
	}
	for name, link := range bookmarks {
		wf.NewItem(name).UID(name).Valid(true).Arg(link)
	}
	log.Println(len(bookmarks))

	// TODO: message doesnt show
	wf.WarnEmpty("No matching items", "Try a different query?")

	wf.SendFeedback()
}

// Return a map of all links in wiki.
func getLinks() {
}

// Parse SUMMARY.md file and return all file paths inside the wiki.
func parseSummary() []string {
	files := make([]string, 0)
	data, err := ioutil.ReadFile("SUMMARY.md")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)
	scanner := bufio.NewScanner(strings.NewReader(string(data)))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		files = append(files, matches[0][2])
	}
	return files
}

// Parse file for links and update bookmarks.
func parseLinks(wikiPath string, filePath string, bookmarks bookmarks) {
	file, err := os.Open(wikiPath + filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)

	scanner := bufio.NewScanner(file)
	reading := false
	for scanner.Scan() {
		if reading == false && strings.HasPrefix(scanner.Text(), "## Links") {
			reading = true
			continue
		}
		if reading == true && strings.HasPrefix(scanner.Text(), "#") {
			break
		} else if reading == true {
			matches := re.FindAllStringSubmatch(scanner.Text(), -1)
			if matches != nil {
				bookmarks[matches[0][1]] = matches[0][2]
			}
		}
	}
}
