package commands

import (
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/nakabonne/giturl/pkg/converter"
)

type runner struct {
	stdout      io.Writer
	scheme      converter.Scheme
	makeOptions func() *converter.Options
}

func (r *runner) run(_ *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no URL given")
	}
	var opts *converter.Options
	if r.makeOptions != nil {
		opts = r.makeOptions()
	}

	res, err := converter.Convert(args[0], r.scheme, opts)
	if err != nil {
		return err
	}
	fmt.Fprintln(r.stdout, res)
	return nil
}
