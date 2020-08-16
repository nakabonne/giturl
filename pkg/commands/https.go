package commands

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/converter"
)

func NewHTTPSCommand(stdout io.Writer) *cobra.Command {
	r := &runner{
		stdout: stdout,
		scheme: converter.SchemeHTTPS,
	}
	cmd := &cobra.Command{
		Use:     "https",
		Short:   "Convert into https syntax",
		Example: "giturl https git@github.com:org/repo.git",
		RunE:    r.run,
	}

	return cmd
}
