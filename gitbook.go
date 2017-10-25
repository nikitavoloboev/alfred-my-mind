// Package gitbook contains helper functions to parse GitBook summaries
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

// getAllLinks returns all links and their names from a given markdown file.
func getAllLinks(fileContent string) {
	m := make(map[string]string)

	// regex to extract link and text attached to link
	re := regexp.MustCompile(`\[([^\]]*)\]\(([^)]*)\)`)

	scanner := bufio.NewScanner(strings.NewReader(fileContent))
	for scanner.Scan() {
		matches := re.FindAllStringSubmatch(scanner.Text(), -1)
		m[matches[0][1]] = matches[0][2]
	}
}

// readFile Reads file.
func readFile(filename string) string {
	b, err := ioutil.ReadFile(filename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}
