package cmd

import (
	"io"
	"os"

	cliflag "github.com/TroyXia/component-base/pkg/cli/flag"
	"github.com/TroyXia/iam/internal/iamctl/util/templates"
	"github.com/spf13/cobra"

	cmdutil "github.com/TroyXia/iam/internal/iamctl/cmd/util"
	"github.com/TroyXia/iam/pkg/cli/genericclioptions"
)

// NewDefaultIAMCtlCommand creates the `iamctl` command with default arguments.
func NewDefaultIAMCommand() *cobra.Command {
	return NewIAMCtlCommand(os.Stdin, os.Stdout, os.Stderr)
}

// NewIAMCtlCommand returns new initialized instance of 'iamctl' root command.
func NewIAMCtlCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   "iamctl",
		Short: "iamctl controls the iam platform",
		Long: templates.LongDesc(`
		iamctl controls the iam platform, is the client side tool for iam platform.

		Find more information at:
			https://github.com/TroyXia/iam/blob/master/docs/guide/en-US/cmd/iamctl/iamctl.md`),
		Run: runHelp,
		// Hook before and after run initialize and write profiles to disk respectively.
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initProfiling()
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return flushProfiling()
		},
	}

	flags := cmds.PersistentFlags()
	flags.SetNormalizeFunc(cliflag.WarnWordSepNormalizeFunc) // Warn for "_" flags

	// Normalize all flags that are coming from other packages or pre-configurations
	// a.k.a. change all "_" to "-". e.g. glog package
	flags.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)

	addProfilingFlags(flags)

	iamConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDeprecatedSecretFlag()
	iamConfigFlags.AddFlags(flags)
	matchVersionIAMConfigFlags := cmdutil.NewMatchVersionFlags(iamConfigFlags)
	matchVersionIAMConfigFlags.AddFlags(cmds.PersistentFlags())

	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
