package main

import (
	"os"

	"github.com/wasilibs/go-jq/internal/runner"
	"github.com/wasilibs/go-jq/internal/wasm"
)

func main() {
	os.Exit(runner.Run("jq", os.Args[1:], wasm.Jq, os.Stdin, os.Stdout, os.Stderr, "."))
}
