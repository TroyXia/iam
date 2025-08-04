package cmd

import (
	"flag"
	cmdutil "github.com/TroyXia/iam/internal/iamctl/cmd/util"
	"github.com/TroyXia/iam/internal/iamctl/cmd/version"
	"github.com/TroyXia/iam/internal/iamctl/util/templates"
	genericapiserver "github.com/TroyXia/iam/internal/pkg/server"
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/marmotedu/iam/pkg/cli/genericclioptions"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io"
	"os"
)

// NewDefaultIAMCtlCommand creates the `iamctl` command with default arguments.
func NewDefaultIAMCtlCommand() *cobra.Command {
	return NewIAMCtlCommand(os.Stdin, os.Stdout, os.Stderr)
}

// NewIAMCtlCommand returns new initialized instance of 'iamctl' root command.
func NewIAMCtlCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	// Parent command to which all subcommands are added.
	cmds := &cobra.Command{
		Use:   "iamctl",
		Short: "iamctl controls the iam platform",
		// TODO: continue explore
		Long: templates.LongDesc(`
		iamctl controls the iam platform, is the client side tool for iam platform.

		Find more information at:
			https://github.com/marmotedu/iam/blob/master/docs/guide/en-US/cmd/iamctl/iamctl.md`),
		Run: runHelp,

		// TODO: continue explore
		// Hook before and after Run initialize and write profiles to disk, respectively.
		PersistentPreRunE: func(*cobra.Command, []string) error {
			return initProfiling()
		},
		PersistentPostRunE: func(*cobra.Command, []string) error {
			return flushProfiling()
		},
	}

	flags := cmds.PersistentFlags()
	flags.SetNormalizeFunc(cliflag.WarnWordSepNormalizeFunc) // Warn for "_" flags

	// Normalize all flags that are coming from other packages or pre-configurations
	// a.k.a. change all "_" to "-". e.g. glog package
	flags.SetNormalizeFunc(cliflag.WordSepNormalizeFunc)

	addProfilingFlags(flags)

	// TODO: continue explore
	iamConfigFlags := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag().WithDeprecatedSecretFlag()
	iamConfigFlags.AddFlags(flags)
	matchVersionIAMConfigFlags := cmdutil.NewMatchVersionFlags(iamConfigFlags)
	matchVersionIAMConfigFlags.AddFlags(cmds.PersistentFlags())

	_ = viper.BindPFlags(cmds.PersistentFlags())
	cobra.OnInitialize(func() {
		genericapiserver.LoadConfig(viper.GetString(genericclioptions.FlagIAMConfig), "iamctl.yaml")
	})
	cmds.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	f := cmdutil.NewFactory(matchVersionIAMConfigFlags)

	// From this point and forward we get warnings on flags that contain "_" separators
	cmds.SetGlobalNormalizationFunc(cliflag.WarnWordSepNormalizeFunc)

	ioStreams := genericclioptions.IOStreams{In: in, Out: out, ErrOut: err}

	cmds.AddCommand(version.NewCmdVersion(f, ioStreams))

	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	_ = cmd.Help()
}
