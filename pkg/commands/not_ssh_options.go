package commands

import (
	"github.com/nakabonne/giturl/pkg/converter"
)

type notSSHOptions struct {
	noUser bool
}

func (o *notSSHOptions) makeOptions() *converter.Options {
	opts := &converter.Options{}
	if o.noUser {
		user := ""
		opts.User = &user
	}
	return opts
}
