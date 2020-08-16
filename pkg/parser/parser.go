package parser

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// Parse parses git url into a URL structure.
// All unparsable strings are recognized as local paths.
func Parse(rawURL string) *url.URL {
	if u, err := parseTransport(rawURL); err == nil {
		return u
	}
	if u, err := parseScp(rawURL); err == nil {
		return u
	}

	return &url.URL{
		Scheme: "file",
		Host:   "",
		Path:   rawURL,
	}
}

var (
	knownSchemes = map[string]interface{}{
		"ssh":     struct{}{},
		"git":     struct{}{},
		"git+ssh": struct{}{},
		"rsync":   struct{}{},
		"http":    struct{}{},
		"https":   struct{}{},
		"file":    struct{}{},
	}
	scpRegex = regexp.MustCompile(`^([a-zA-Z0-9_]+@)?([a-zA-Z0-9._-]+):(.*)$`)
)

// Return a structured URL only when scheme is a known Git transport.
func parseTransport(rawURL string) (*url.URL, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse git url: %w", err)
	}
	if _, ok := knownSchemes[u.Scheme]; !ok {
		return nil, fmt.Errorf("unknown scheme %q", u.Scheme)
	}
	return u, nil
}

// Return a structured URL only when the rawURL is an SCP-like URL.
func parseScp(rawURL string) (*url.URL, error) {
	match := scpRegex.FindAllStringSubmatch(rawURL, -1)
	if len(match) == 0 {
		return nil, fmt.Errorf("%q is not an SCP-like URL", rawURL)
	}
	m := match[0]
	var (
		user     = strings.TrimRight(m[1], "@")
		host     = m[2]
		path     = m[3]
		userinfo *url.Userinfo
	)
	if user != "" {
		userinfo = url.User(user)
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return &url.URL{
		Scheme: "ssh",
		User:   userinfo,
		Host:   host,
		Path:   path,
	}, nil
}
