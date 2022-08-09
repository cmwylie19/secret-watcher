package main

import (
	"fmt"
	"os"

	"github.com/cmwylie19/secret-watcher/internal/cmd"
)

func main() {
	if err := cmd.GetRootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
