# Contributing to Alfred My Mind

Thank you for taking the time to contribute! ♥️

You can:

- Submit [bug reports or feature requests](../../issues/new/choose)
- Fix [open issues](../../issues) and [create PRs](https://help.github.com/en/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request)
- Improve documentation, the code and more! These are just some ideas.

## Dependencies

The workflow is written in [Go](https://golang.org/) and uses [AwGo](https://github.com/deanishe/awgo#readme) library for all Alfred related things.

It uses [modd](https://github.com/cortesi/modd#readme) and [Alfred command](https://godoc.org/github.com/jason0x43/go-alfred/alfred) to help with its development.

To run the project:

1. Clone the repo (`git clone https://github.com/nikitavoloboev/alfred-my-mind`)
2. Run `alfred link` (makes symbolic link of [`workflow`](workflow) directory)
3. Run `modd` (starts a process that automatically builds the workflow with `alfred build` on any changes you make to `.go` files, this builds and places a binary inside [`workflow`](workflow) directory.)
4. Make changes to code or modify Alfred objects to do what you want! Open debugger in Alfred or run the workflow with `workflow:log` passed in as argument to see the logs Alfred produces.

![](https://i.imgur.com/FFYOecx.png)

Have fun! 🚀

You can also read [how to write Alfred workflows in Go](https://medium.com/@nikitavoloboev/writing-alfred-workflows-in-go-2a44f62dc432).

## Opening PRs

Please search [existing issues](../../issues/) and [pull requests](../../pulls/) to check if somebody else is already working on the issue.

Don't be afraid to ask questions, open issues for discussion or opening [Draft PRs](https://github.blog/2019-02-14-introducing-draft-pull-requests/) with code.
