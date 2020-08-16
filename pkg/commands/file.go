package commands

import (
	"io"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/converter"
)

func NewFileCommand(stdout io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "file",
		Short:   "Convert into file syntax",
		Example: "giturl file git@github.com:org/repo.git",
	}

	r := &runner{
		stdout: stdout,
		scheme: converter.SchemeFile,
	}
	cmd.RunE = r.run

	return cmd
}
