package commands

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/converter"
)

func NewGitCommand(stdout io.Writer) *cobra.Command {
	o := &httpsOptions{}
	cmd := &cobra.Command{
		Use:     "git",
		Short:   "Convert into git syntax",
		Example: "giturl git --no-user git@github.com:org/repo.git",
	}
	cmd.Flags().BoolVarP(&o.noUser, "no-user", "n", o.noUser, "prune user from the given URL")

	r := &runner{
		stdout:      stdout,
		scheme:      converter.SchemeGit,
		makeOptions: o.makeOptions,
	}
	cmd.RunE = r.run

	return cmd
}
