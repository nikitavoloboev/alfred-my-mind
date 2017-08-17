package main

import (
	// "log"

	"git.deanishe.net/deanishe/awgo"
	"git.deanishe.net/deanishe/awgo/update"
	"gopkg.in/alecthomas/kingpin.v2"
)

// name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var (
	app   *kingpin.Application
	query string // script options
	repo  = "nikitavoloboev/alfred-my-mind"
	wf    *aw.Workflow
)

func init() {
	wf = aw.New(update.GitHub(repo))

	app = kingpin.New("query-maps", "search and query interactive maps")

}

func parseMaps() {

}

func run() {

	// var query string
	// args := wf.Args()

	// if len(args) > 0 {
	// 	query = args[0]
	// }

	// if query != "" {
	// 	// sort results
	// 	res := wf.Filter(query)
	// 	log.Printf("%d results match `%s`", len(res), query)

	// 	for i, r := range res {
	// 		log.Printf("%02d. score=%0.1f sortkey=%s", i+1, r.Score, wf.Feedback.SortKey(i))
	// 	}
	// }

	// wf.SendFeedback()

}

func main() {
	wf.Run(run)
}
