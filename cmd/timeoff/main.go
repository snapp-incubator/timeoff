package main

import (
	"os"

	"github.com/snapp-incubator/timeoff/internal/app/timeoff/cmd"
	_ "go.uber.org/automaxprocs"
)

const exitFailure = 1

func main() {
	root := cmd.NewRootCommand()

	if root != nil {
		if err := root.Execute(); err != nil {
			os.Exit(exitFailure)
		}
	}
}
