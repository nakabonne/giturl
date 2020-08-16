package parser

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		rawURL  string
		wantURL *url.URL
	}{
		{
			name:   "SCP-like URL with user",
			rawURL: "user@host.xz:path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "ssh",
				User:   url.User("user"),
				Host:   "host.xz",
				Path:   "/path/to/repo.git/",
			},
		},
		{
			name:   "SCP-like URL without user",
			rawURL: "host.xz:path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "ssh",
				User:   nil,
				Host:   "host.xz",
				Path:   "/path/to/repo.git/",
			},
		},
		{
			name:   "SCP-like URL with prefix `/`",
			rawURL: "host.xz:/path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "ssh",
				User:   nil,
				Host:   "host.xz",
				Path:   "/path/to/repo.git/",
			},
		},
		{
			name:   "ssh with user",
			rawURL: "ssh://user@host.xz/path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "ssh",
				User:   url.User("user"),
				Host:   "host.xz",
				Path:   "/path/to/repo.git/",
			},
		},
		{
			name:   "ssh with user with port",
			rawURL: "ssh://user@host.xz:1234/path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "ssh",
				User:   url.User("user"),
				Host:   "host.xz:1234",
				Path:   "/path/to/repo.git/",
			},
		},
		{
			name:   "git+ssh",
			rawURL: "git+ssh://host.xz/path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "git+ssh",
				User:   nil,
				Host:   "host.xz",
				Path:   "/path/to/repo.git/",
			},
		},
		{
			name:   "file scheme",
			rawURL: "file:///path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "file",
				User:   nil,
				Host:   "",
				Path:   "/path/to/repo.git/",
			},
		},
		{
			name:   "rsync + ssh",
			rawURL: "rsync://host.xz/path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "rsync",
				User:   nil,
				Host:   "host.xz",
				Path:   "/path/to/repo.git/",
			},
		},
		{
			name:   "git scheme",
			rawURL: "git://host.xz/path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "git",
				User:   nil,
				Host:   "host.xz",
				Path:   "/path/to/repo.git/",
			},
		},
		{
			name:   "http scheme",
			rawURL: "http://host.xz/path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "http",
				User:   nil,
				Host:   "host.xz",
				Path:   "/path/to/repo.git/",
			},
		},
		{
			name:   "https scheme",
			rawURL: "https://host.xz/path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "https",
				User:   nil,
				Host:   "host.xz",
				Path:   "/path/to/repo.git/",
			},
		},
		{
			name:   "local repository without file scheme",
			rawURL: "/path/to/repo.git/",
			wantURL: &url.URL{
				Scheme: "file",
				User:   nil,
				Host:   "",
				Path:   "/path/to/repo.git/",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Parse(tt.rawURL)
			assert.Equal(t, tt.wantURL, got)
		})
	}
}
