package main

import (
	"github.com/alecthomas/kingpin"
	"github.com/deanishe/awgo"
	"github.com/deanishe/awgo/update"
)

var (
	// Icons
	updateAvailable = &aw.Icon{Value: "icons/update-available.png"}

	// Kingpin and script options
	app *kingpin.Application

	// Application commands
	searchWikiCmd *kingpin.CmdClause
	updateCmd     *kingpin.CmdClause

	query string

	repo = "nikitavoloboev/alfred-my-mind"

	// Workflow stuff
	wf *aw.Workflow
)

// Mostly sets up kingpin commands
func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func run() {
	showUpdateStatus()
	links := parseSummaryFile()
	for _, link := range links {
		wf.NewItem(link.name).UID(link.uid).Valid(true).Arg(link.uid)
	}

	args := wf.Args()
	var searchQuery string
	if len(args) > 0 {
		searchQuery = args[0]
	}

	if searchQuery == "" {
		wf.WarnEmpty("No Keyboard Maestro macros found", "It seems that you haven't created any macro yet.")
	} else {
		wf.Filter(searchQuery)
		wf.WarnEmpty("No Keyboard Maestro macros found that matched your query", "Try a different query?")
	}

	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
