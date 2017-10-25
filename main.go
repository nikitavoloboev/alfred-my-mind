package main

import (
	"log"
	"os/exec"

	"git.deanishe.net/deanishe/awgo"
	"git.deanishe.net/deanishe/awgo/update"
	"github.com/docopt/docopt-go"
)

// TODO: don't use docopt

// Name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var usage = `alfred-my-mind [search|check] [<query>]

Access notes, wiki and more

Usage:
	alfred-web-searches search [<query>]
	alfred-web-searches check
    alfred-web-searches -h

Options:
    -h, --help    Show this message and exit.
`

var (
	// icons
	iconAvailable = &aw.Icon{Value: "icons/update.png"}

	repo = "nikitavoloboev/alfred-web-searches"
	wf   *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo))
}

func run() {
	// Pass wf.Args() to docopt because update logic relies on
	// AwGo's magic actions.
	args, _ := docopt.Parse(usage, wf.Args(), true, wf.Version(), false, true)

	// alternate action: get available releases from remote
	if args["check"] != false {
		wf.TextErrors = true
		log.Println("checking for updates...")
		if err := wf.CheckForUpdate(); err != nil {
			wf.FatalError(err)
		}
		return
	}

	var query string
	if args["<query>"] != nil {
		query = args["<query>"].(string)
	}

	log.Printf("query=%s", query)

	// call self with "check" command if an update is due and a
	// check job isn't already running.
	if wf.UpdateCheckDue() && !aw.IsRunning(updateJobName) {
		log.Println("running update check in background...")
		cmd := exec.Command("./alfred-web-searches", "check")
		if err := aw.RunInBackground(updateJobName, cmd); err != nil {
			log.Printf("error starting update check: %s", err)
		}
	}

	if query == "" { // Only show update status if query is empty
		// Send update status to Alfred
		if wf.UpdateAvailable() {
			wf.NewItem("update available!").
				Subtitle("↩ to install").
				Autocomplete("workflow:update").
				Valid(false).
				Icon(iconAvailable)
		}
	}

	parseSummary()

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
