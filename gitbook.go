// Package gitbook contains helper functions to parse GitBook summaries.
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

// popLine removes lines from file.
func popLine(f *os.File) ([]byte, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(buf, f)
	if err != nil {
		return nil, err
	}
	line, err := buf.ReadString('\n')
	if err != nil && err != io.EOF {
		return nil, err
	}

	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	nw, err := io.Copy(f, buf)
	if err != nil {
		return nil, err
	}
	err = f.Truncate(nw)
	if err != nil {
		return nil, err
	}
	err = f.Sync()
	if err != nil {
		return nil, err
	}

	_, err = f.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil, err
	}
	return []byte(line), nil
}

// getSummary downloads GitBook summary from GitHub.
func getSummary() ([]byte, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/nikitavoloboev/knowledge/master/SUMMARY.md")
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	return buf.Bytes(resp.Body), nil

	// f, err := os.Create("SUMMARY.md")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer f.Close()

	// io.Copy(f, resp.Body)

}

func removeFirstLine(fname string) {
	f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	_, err = popLine(f)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func doSearchWiki() error {
	showUpdateStatus()

	if err := wf.Session.LoadOrStore("wiki", 60*time.Minute, getSummary()) {
		return nil, err
	}

	log.Printf("query=%s", query)

	parseSummaryFile()

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()

	return nil
}

// parseSummaryFile parses GitBook Summary.md file
func parseSummaryFile() {
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
