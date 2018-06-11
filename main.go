package main

import (
	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	// Icons
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}

	query string

	repo = "nikitavoloboev/alfred-my-mind"

	wf *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func run() {
	showUpdateStatus()
	links := parseSummaryFile()
	for _, link := range links {
		wf.NewItem(link.name).UID(link.uid).Valid(true).Arg(link.url)
	}

	args := wf.Args()
	if len(args) > 0 {
		query = args[0]
	}

	// FIXME:
	wf.WarnEmpty("No matching items", "Try a different query?")

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
