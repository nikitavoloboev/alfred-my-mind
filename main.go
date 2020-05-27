package main

import (
	"flag"

	aw "github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	iconUpdate = &aw.Icon{Value: "icons/update-available.png"}

	query string

	repo = "nikitavoloboev/alfred-my-mind"

	wf *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func run() {
	wiki := flag.Bool("wiki", false, "Search wiki")
	links := flag.Bool("links", false, "Search links in wiki")
	flag.Parse()
	if *wiki {
		searchWiki()
		return
	} else if *links {
		// TODO: read from env var
		// TODO: show description as subtitle
		searchLinks("/Users/nikivi/Dropbox/Write/knowledge/")
		return
	}
}

func main() {
	wf.Run(run)
}
