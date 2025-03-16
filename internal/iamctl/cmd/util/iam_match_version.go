package util

import (
	"sync"

	"github.com/TroyXia/iam/pkg/cli/genericclioptions"
	"github.com/spf13/pflag"
)

const (
	flagMatchBinaryVersion = "match-server-version"
)

// MatchVersionFlags is for setting the "match server version" function.
type MatchVersionFlags struct {
	Delegate genericclioptions.RESTClientGetter

	RequireMatchedServerVersion bool
	checkServerVersion          sync.Once
	matchesServerVersionErr     error
}

func (f *MatchVersionFlags) AddFlags(flags *pflag.FlagSet) {
	flags.BoolVar(
		&f.RequireMatchedServerVersion,
		flagMatchBinaryVersion,
		f.RequireMatchedServerVersion,
		"Require server version to match client version",
	)
}

func NewMatchVersionFlags(delegate genericclioptions.RESTClientGetter) *MatchVersionFlags {
	return &MatchVersionFlags{
		Delegate: delegate,
	}
}
