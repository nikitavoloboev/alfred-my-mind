package main

import (
	// "encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "os"

	"git.deanishe.net/deanishe/awgo"
	"git.deanishe.net/deanishe/awgo/update"
	"gopkg.in/alecthomas/kingpin.v2"
)

// name of the background job that checks for updates
const updateJobName = "checkForUpdate"

var (
	app *kingpin.Application

	// app commands
	updateMapsCmd, downloadMapsCmd, parseMapsCmd *kingpin.CmdClause

	mindnodeUrls []string // contains links to JSON of maps

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
	// parseMapsCmd = app.Command("parse", "parses maps")

	app.DefaultEnvars()
}

// _actions
func parseMaps(file string) error {
	return nil
}

// updateMaps
func updateMaps() error {
	wf.NewItem("hello")
	wf.SendFeedback()
	return nil
}

// parseUrls parses urls from maps.json
func parseUrls(filename string) {
	// f, err := os.Open(filename)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}

// downloadMaps parses maps.json and downloads maps specified there to maps directory
func downloadMaps(filename string) error {
	parseUrls(filename)

	// f, err := os.Open(filename)
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	// _, err = os.Stat("test.txt")
	// if os.IsNotExist(err) { // file does not exist
	// 	log.Fatal(err)
	// }

	// log.Printf("created file")
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
		err = downloadMaps("maps.json")
	case parseMapsCmd.FullCommand():
		err = parseMaps("maps.json")
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
