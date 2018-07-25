# Contributing to Alfred My Mind
There are many ways you can contribute. You can:
- Make suggestions and file bugs in [Issues](../../issues/).
- Fix issues and add features with [Pull Requests](../../pulls/).

## Dependencies
The workflow is written in [Go](https://golang.org/) and uses [AwGo](https://github.com/deanishe/awgo) library for all Alfred related things.

## Developing the workflow
If you want to develop the workflow. It is best to use [Alfred CLI tool](https://godoc.org/github.com/jason0x43/go-alfred/alfred) which you can install by running:

`go install github.com/jason0x43/go-alfred/alfred`

You can then clone this repository and run: `alfred link` inside it. This will make a symbolic link of the [`workflow`](workflow) directory.

You can then make changes to the code and after run `alfred build` to build the go binary to `workflow` directory. Which you can then call from Alfred's [Script Filters](https://www.alfredapp.com/help/workflows/inputs/script-filter/).

You can also read [this article](https://medium.com/@nikitavoloboev/writing-alfred-workflows-in-go-2a44f62dc432) on how to write your own Alfred workflows in Go.

## Submitting a Pull Request
Please go through [existing issues](../../issues/) and [pull requests](../../pulls/) to check if somebody else is already working on the issue.

#### Thank you for taking the time to contribute! 💜