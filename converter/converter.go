package converter

import "github.com/nakabonne/giturl/parser"

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
func Convert(rawURL string, scheme Scheme) string {
	parser.Parse(rawURL)
}
