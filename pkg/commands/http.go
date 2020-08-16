package commands

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/converter"
)

func NewHTTPCommand(stdout io.Writer) *cobra.Command {
	r := &runner{
		stdout: stdout,
		scheme: converter.SchemeHTTP,
	}
	cmd := &cobra.Command{
		Use:   "http",
		Short: "Convert into http syntax",
		RunE:  r.run,
	}

	return cmd
}
