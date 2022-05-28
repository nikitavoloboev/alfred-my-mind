//go:build mage

package main

// Uses https://github.com/jason0x43/go-alfred to build workflow for use
// TODO: don't rely on CLI for this, export it as part of go lib
func Build() {
	// run `alfred build`
}

// Symlink `workflow` dir to Alfred
func LinkWorkflow() {
	// run `alfred link`
}
