package main

import (
	"github.com/goyek/x/boot"
	"github.com/wasilibs/tools/tasks"
)

func main() {
	tasks.Define(tasks.Params{
		LibraryName: "jq",
		LibraryRepo: "jqlang/jq",
		GoReleaser:  true,
	})
	boot.Main()
}
