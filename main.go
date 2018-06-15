package main

import (
	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}

	query string

	repo = "nikitavoloboev/alfred-my-mind"

	wf *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func run() {
	searchWiki()
}

func main() {
	wf.Run(run)
}
