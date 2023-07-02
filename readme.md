# Alfred My Mind [![Workflows](https://img.shields.io/badge/-more%20workflows-0a0a0a.svg?style=flat&colorA=0a0a0a)](https://github.com/learn-anything/alfred-workflows) [![Thanks](http://bit.ly/saythankss)](https://github.com/sponsors/nikitavoloboev)

> [Alfred](https://www.alfredapp.com/) workflow to search through my notes and bookmarks

<img src="https://i.imgur.com/ZbpobeM.png" width="500" alt="img">

This workflow lets you search through entirety of [my personal wiki](https://wiki.nikiv.dev/). There's over 100,000 lines of markdown with many links inside. This [video](https://www.youtube.com/watch?v=m5aFsVVknPU) showcases the workflow in action.

## Install

Download workflow from [GitHub releases](../../releases/latest).

See [here](https://github.com/deanishe/awgo/wiki/Catalina) for instructions on fixing permissions in macOS refusing to run Go binary.

## Setup

The workflow is written in [Go](https://golang.org/) and uses [AwGo](https://github.com/deanishe/awgo) library for all Alfred related things.

It uses [modd](https://github.com/cortesi/modd) and [Alfred command](https://godoc.org/github.com/jason0x43/go-alfred/alfred) to ease its development.

## Run

1. Run `alfred link` (makes symbolic link of [`workflow`](workflow) directory)
2. Run `modd` (starts a process that automatically builds the workflow with `alfred build` on any changes you make to `.go` files, this builds and places a binary inside [`workflow`](workflow) directory.)
3. Make changes to code or modify Alfred objects to do what you want! Open debugger in Alfred or run the workflow with `workflow:log` passed in as argument to see the logs Alfred produces.

![](https://i.imgur.com/FFYOecx.png)

## Thank you

You can support me on [GitHub](https://github.com/sponsors/nikitavoloboev) or look into [other projects](https://nikiv.dev/projects) I shared.

I also have [personal Discord](https://discord.gg/f8YHjyrX3h) you can join for more indepth discussions.

[![MIT](http://bit.ly/mitbadge)](https://choosealicense.com/licenses/mit/) [![Twitter](http://bit.ly/nikitatweet)](https://twitter.com/nikitavoloboev)
