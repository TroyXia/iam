package main

import (
	"os"

	"github.com/TroyXia/iam/internal/iamctl/cmd"
)

func main() {
	command := cmd.NewDefaultIAMCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
