package converter

import (
	"fmt"
	"net/url"

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

// Options lets you control the output URLs in detail.
type Options struct {
	// Override the user.
	User string
	// Prune user from the given URL.
	PruneUser bool
	// Return SCP-like URL instead of the general ssh syntax.
	ScpLike bool
}

// Convert will attempt to convert to the given scheme.
// Returns an error if the conversion isn't possible.
func Convert(rawURL string, scheme Scheme, opts Options) (string, error) {
	u := parser.Parse(rawURL)
	if opts.User != "" {
		u.User = url.User(opts.User)
	}
	var res string
	switch scheme {
	case SchemeSSH:
		if opts.ScpLike {
			res = fmt.Sprintf("%s:%s", u.Host, u.Path)
			if u.User != nil {
				res = fmt.Sprintf("%s@%s:%s", u.User, u.Host, u.Path)
			}
			break
		}
		res = fmt.Sprintf("ssh://%s@%s/%s", u.User, u.Host, u.Path)
		if u.User != nil {
			res = fmt.Sprintf("ssh://%s/%s", u.Host, u.Path)
		}
	case SchemeHTTP:
		res = fmt.Sprintf("http://%s/%s", u.Host, u.Path)
		if u.User != nil {
			res = fmt.Sprintf("http://%s@%s/%s", u.User, u.Host, u.Path)
		}
	case SchemeHTTPS:
		res = fmt.Sprintf("https://%s/%s", u.Host, u.Path)
		if u.User != nil {
			res = fmt.Sprintf("https://%s@%s/%s", u.User, u.Host, u.Path)
		}
	default:
		return "", fmt.Errorf("unsupported scheme %q", scheme)
	}
	return res, nil
}
