package main

import (
	"flag"
	"fmt"

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
	inside := flag.Bool("inside", false, "Search links inside files")
	flag.Parse()

	if *wiki {
		searchWiki()
		return
	}

	if *links {
		// TODO: read from env var
		// TODO: show description as subtitle
		searchLinks("/Users/nikivi/Dropbox/Write/knowledge/")
		return
	}

	if *inside {
		fmt.Println("hi")
		return
	}
}

func main() {
	wf.Run(run)
}
