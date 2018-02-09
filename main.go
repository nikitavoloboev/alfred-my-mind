package main

import (
	"fmt"

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
	searchCmd *kingpin.CmdClause
	updateCmd *kingpin.CmdClause

	query string

	repo = "nikitavoloboev/alfred-my-mind"

	// Workflow stuff
	wf *aw.Workflow
)

// Mostly sets up kingpin commands
func init() {
	wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))

	app = kingpin.New("ask", "Search my notes.")

	// Update command
	updateCmd = app.Command("update", "Check for new version.").Alias("u")

	// Commands using query
	searchCmd = app.Command("search", "Search websites.").Alias("s")

	// Common options
	for _, cmd := range []*kingpin.CmdClause{
		searchCmd,
	} {
		cmd.Flag("query", "Search query.").Short('q').StringVar(&query)
	}
}

func run() {
	var err error

	cmd, err := app.Parse(wf.Args())
	if err != nil {
		wf.FatalError(err)
	}

	switch cmd {
	case searchCmd.FullCommand():
		err = doSearch()
	case updateCmd.FullCommand():
		err = doUpdate()
	default:
		err = fmt.Errorf("Uknown command: %s", cmd)
	}

	// Check for update
	if err == nil && cmd != updateCmd.FullCommand() {
		err = checkForUpdate()
	}

	if err != nil {
		wf.FatalError(err)
	}
}

func main() {
	wf.Run(run)
}
