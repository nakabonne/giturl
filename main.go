package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime"

	flag "github.com/spf13/pflag"

	"github.com/nakabonne/giturl/pkg/converter"
)

var (
	flagSet = flag.NewFlagSet("giturl", flag.ContinueOnError)

	usage = func() {
		fmt.Fprintln(os.Stderr, "usage: giturl [<flag> ...] <Git URL> ")
		flagSet.PrintDefaults()
	}

	// Automatically populated by goreleaser during build
	version = "unversioned"
	commit  = "?"
	date    = "?"
)

type app struct {
	scheme  string
	user    string
	noUser  bool
	scpLike bool
	version bool
	stdout  io.Writer
	stderr  io.Writer
}

func main() {
	a := &app{
		scheme: "ssh",
		stdout: os.Stdout,
		stderr: os.Stderr,
	}
	flagSet.StringVarP(&a.scheme, "scheme", "s", a.scheme, "convert to the given schema: ssh|http|https|git|file")
	flagSet.StringVar(&a.user, "user", "", "set user")
	flagSet.BoolVar(&a.noUser, "no-user", false, "prune user from the given URL")
	flagSet.BoolVarP(&a.scpLike, "scp-like", "S", false, "emit scp-like syntax (available only when --schema=ssh)")
	flagSet.BoolVarP(&a.version, "version", "v", false, "print the current version")
	flagSet.Usage = usage

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Fprintln(a.stderr, err)
		}
		return
	}

	os.Exit(a.run(flagSet.Args()))
}

func (a *app) run(args []string) int {
	if a.version {
		fmt.Fprintf(a.stderr, "version=%s, commit=%s, buildDate=%s, os=%s, arch=%s\n", version, commit, date, runtime.GOOS, runtime.GOARCH)
		return 0
	}

	if len(args) == 0 {
		fmt.Fprintln(a.stderr, "No URL given")
		usage()
		return 1
	}
	if a.scheme == "" {
		a.scheme = "ssh"
	}
	user := a.user
	if a.noUser {
		user = ""
	}
	opts := converter.Options{
		User:    &user,
		ScpLike: a.scpLike,
	}
	res, err := converter.Convert(args[0], converter.Scheme(a.scheme), opts)
	if err != nil {
		fmt.Fprintln(a.stderr, err)
		usage()
		return 1
	}
	fmt.Fprintf(a.stdout, "%s\n", res)
	return 0
}
