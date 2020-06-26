# Contributing

Thank you for taking the time to contribute! ♥️ You can:

- Submit [bug reports or feature requests](../../issues/new/choose). Contribute to discussions. Fix [open issues](../../issues).
- Improve docs, the code and more! Any idea is welcome.

## Run project

The workflow is written in [Go](https://golang.org/) and uses [AwGo](https://github.com/deanishe/awgo) library for all Alfred related things.

It uses [modd](https://github.com/cortesi/modd) and [Alfred command](https://godoc.org/github.com/jason0x43/go-alfred/alfred) to ease its development.

1. Clone repo
2. Run `alfred link` (makes symbolic link of [`workflow`](workflow) directory)
3. Run `modd` (starts a process that automatically builds the workflow with `alfred build` on any changes you make to `.go` files, this builds and places a binary inside [`workflow`](workflow) directory.)
4. Make changes to code or modify Alfred objects to do what you want! Open debugger in Alfred or run the workflow with `workflow:log` passed in as argument to see the logs Alfred produces.

![](https://i.imgur.com/FFYOecx.png)
