package commands

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/converter"
)

func NewHTTPSCommand(stdout io.Writer) *cobra.Command {
	o := &notSSHOptions{}
	cmd := &cobra.Command{
		Use:     "https",
		Short:   "Convert into https syntax",
		Example: "giturl https --no-user git@github.com:org/repo.git",
	}
	cmd.Flags().BoolVarP(&o.noUser, "no-user", "n", o.noUser, "prune user from the given URL")

	r := &runner{
		stdout:      stdout,
		scheme:      converter.SchemeHTTPS,
		makeOptions: o.makeOptions,
	}
	cmd.RunE = r.run

	return cmd
}
