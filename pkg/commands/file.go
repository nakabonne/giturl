package commands

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/converter"
)

func NewFileCommand(stdout io.Writer) *cobra.Command {
	r := &runner{
		stdout: stdout,
		scheme: converter.SchemeFile,
	}
	cmd := &cobra.Command{
		Use:   "file",
		Short: "Convert into file syntax",
		RunE:  r.run,
	}

	return cmd
}
