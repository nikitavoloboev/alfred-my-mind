# Contributing to Alfred My Mind
There are many ways you can contribute. You can:
- Make suggestions and file bugs in [Issues](../../issues/).
- Fix issues and add features with [Pull Requests](../../pulls/).

## Dependencies
The workflow is written in [Go](https://golang.org/) and uses [AwGo](https://github.com/deanishe/awgo#readme) library for all Alfred related things.

It also uses [modd](https://github.com/cortesi/modd#readme) and [Alfred command](https://godoc.org/github.com/jason0x43/go-alfred/alfred) to help with its development.

## Developing the workflow
1. Clone this repo and run: `alfred link` inside it. This will make a symbolic link of the [`workflow`](workflow) directory.
2. Running `modd` will start a process that will automatically build the workflow with `alfred build` on any changes you make to `.go` files. This builds and places a binary inside [`workflow`](workflow) directory.
3. Make changes to the code or modify Alfred objects to do what you want! Open debugger in Alfred or run the workflow with `workflow:log` passed in as argument to see the logs Alfred produces.

![](https://i.imgur.com/FZ91Qkc.png)

You can also read [this article](https://medium.com/@nikitavoloboev/writing-alfred-workflows-in-go-2a44f62dc432) on how to write Alfred workflows in Go.

## Sending PRs
1. Fork repo
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Create new Pull Request

For bigger code changes, it's best to first discuss what you want to add in an issue.

## Submitting a Pull Request
Please go through [existing issues](../../issues/) and [pull requests](../../pulls/) to check if somebody else is already working on the issue.

#### Thank you for taking the time to contribute! 💜