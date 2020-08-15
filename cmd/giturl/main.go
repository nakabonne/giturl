package main

import (
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	flagSet = flag.NewFlagSet("giturl", flag.ContinueOnError)

	usage = func() {
		fmt.Fprintln(os.Stderr, "usage: giturl [<flag> ...] <Git URL> ")
		flagSet.PrintDefaults()
	}
)

type app struct {
	scheme  string
	user    string
	noUser  bool
	scpLike bool
	verbose bool
	stdout  io.Writer
	stderr  io.Writer
}

func main() {
	a := &app{
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	flagSet.StringVarP(&a.scheme, "scheme", "s", "ssh", "convert to the given schema: ssh|http|https|git|file")
	flagSet.StringVarP(&a.scheme, "user", "u", "", "set user")
	flagSet.BoolVar(&a.noUser, "no-user", false, "prune user from the given URL")
	flagSet.BoolVar(&a.scpLike, "scp-like", false, "emit scp-like syntax (available only when --schema=ssh)")
	flagSet.Usage = usage

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintln(a.stderr, err)
		}
		return
	}

	os.Exit(a.run(flagSet.Args()))
}

func (a *app) run(args []string) int {
	return 0
}
