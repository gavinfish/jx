package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jenkins-x/jx/pkg/jx/cmd/templates"
)

// EditOptions contains the CLI options
type EditOptions struct {
	*CommonOptions
}

var (
	exit_long = templates.LongDesc(`
		Edit a resource

`)

	exit_example = templates.Examples(`
		# Lets edit the staging Environment
		jx edit env staging
	`)
)

// NewCmdEdit creates the edit command
func NewCmdEdit(commonOpts *CommonOptions) *cobra.Command {
	options := &EditOptions{
		commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "edit [flags]",
		Short:   "Edit a resource",
		Long:    exit_long,
		Example: exit_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			CheckErr(err)
		},
		SuggestFor: []string{"modify"},
	}

	cmd.AddCommand(NewCmdCreateBranchPattern(commonOpts))
	cmd.AddCommand(NewCmdEditAddon(commonOpts))
	cmd.AddCommand(NewCmdEditAppJenkinsPlugins(commonOpts))
	cmd.AddCommand(NewCmdEditBuildpack(commonOpts))
	cmd.AddCommand(NewCmdEditConfig(commonOpts))
	cmd.AddCommand(NewCmdEditDeployKind(commonOpts))
	cmd.AddCommand(NewCmdEditEnv(commonOpts))
	cmd.AddCommand(NewCmdEditHelmBin(commonOpts))
	cmd.AddCommand(NewCmdEditStorage(commonOpts))
	cmd.AddCommand(NewCmdEditUserRole(commonOpts))
	cmd.AddCommand(NewCmdEditExtensionsRepository(commonOpts))

	err := addTeamSettingsCommandsFromTags(cmd, options)
	CheckErr(err)

	return cmd
}

// Run implements this command
func (o *EditOptions) Run() error {
	return o.Cmd.Help()
}
