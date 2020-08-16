package commands

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/converter"
)

func NewFileCommand(stdout io.Writer) *cobra.Command {
	o := &notSSHOptions{}
	cmd := &cobra.Command{
		Use:     "file",
		Short:   "Convert into file syntax",
		Example: "giturl file --no-user git@github.com:org/repo.git",
	}
	cmd.Flags().BoolVarP(&o.noUser, "no-user", "n", o.noUser, "prune user from the given URL")

	r := &runner{
		stdout:      stdout,
		scheme:      converter.SchemeFile,
		makeOptions: o.makeOptions,
	}
	cmd.RunE = r.run

	return cmd
}
