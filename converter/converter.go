package converter

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/nakabonne/giturl/parser"
)

type Scheme string

const (
	SchemeSSH   = "ssh"
	SchemeHTTP  = "http"
	SchemeHTTPS = "https"
	SchemeGit   = "git"
	SchemeFile  = "file"
)

// Options provides an ability to customize the output URLs in detail.
type Options struct {
	// Override the user if non-empty string set.
	// Prune user from the given URL if empty string set.
	User *string
	// Return SCP-like URL instead of the general ssh syntax.
	ScpLike bool
}

// Convert will attempt to convert to the given scheme.
func Convert(rawURL string, scheme Scheme, opts Options) (string, error) {
	u := parser.Parse(rawURL)
	if opts.User != nil {
		if *opts.User == "" {
			u.User = nil
		} else {
			u.User = url.User(*opts.User)
		}
	}

	switch scheme {
	case SchemeSSH:
		if opts.ScpLike {
			u.Path = strings.TrimLeft(u.Path, "/")
			res := fmt.Sprintf("%s:%s", u.Host, u.Path)
			if u.User != nil {
				res = fmt.Sprintf("%s@%s:%s", u.User, u.Host, u.Path)
			}
			return res, nil
		}
		u.Scheme = "ssh"
		return u.String(), nil
	case SchemeHTTP:
		u.Scheme = "http"
		return u.String(), nil
	case SchemeHTTPS:
		u.Scheme = "https"
		return u.String(), nil
	case SchemeGit:
		u.Scheme = "git"
		return u.String(), nil
	case SchemeFile:
		return fmt.Sprintf("file://%s", u.Path), nil
	default:
		return "", fmt.Errorf("unsupported scheme %q", scheme)
	}
}
