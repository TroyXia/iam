package main

import (
	"github.com/TroyXia/iam/internal/iamctl/cmd"
	"os"
)

func main() {
	command := cmd.NewDefaultIAMCtlCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
