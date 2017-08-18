package main

import (
	"fmt"
	// "log"

	"git.deanishe.net/deanishe/awgo"
	"git.deanishe.net/deanishe/awgo/update"
	"gopkg.in/alecthomas/kingpin.v2"
)

// name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var (
	app *kingpin.Application

	// app commands
	updateMapsCmd, downloadMapsCmd *kingpin.CmdClause

	query string // script options
	repo  = "nikitavoloboev/alfred-my-mind"
	wf    *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo))

	// set up kingpin
	app = kingpin.New("alfred-my-mind", "search and query my personal interactive maps")
	app.HelpFlag.Short('h')
	app.Version(wf.Version())

	updateMapsCmd = app.Command("update", "updates maps")
	downloadMapsCmd = app.Command("download", "downloads maps")

	app.DefaultEnvars()
}

// _actions
// parses maps.json and downloads maps to maps dir
func parseMaps() error {
	wf.NewItem("hello")
	wf.SendFeedback()
	return nil
}

func updateMaps() error {
	wf.NewItem("hello")
	wf.SendFeedback()
	return nil
}

func downloadMaps() error {
	wf.NewItem("here")
	wf.SendFeedback()
	return nil
}

func run() {

	var err error

	cmd, err := app.Parse(wf.Args())
	if err != nil {
		wf.FatalError(err)
	}

	switch cmd {
	case updateMapsCmd.FullCommand():
		err = updateMaps()
	case downloadMapsCmd.FullCommand():
		err = downloadMaps()
	default:
		err = fmt.Errorf("unknown command : %s", cmd)
	}

	if err != nil {
		wf.FatalError(err)
	}
}

func main() {
	wf.Run(run)
}
