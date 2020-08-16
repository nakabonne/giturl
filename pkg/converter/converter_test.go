package converter

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		rawURL  string
		scheme  Scheme
		opts    *Options
		want    string
		wantErr bool
	}{
		// The cases for converting https.
		{
			name:    "convert https into ssh",
			rawURL:  "https://host.xz/path/to/repo.git",
			scheme:  "ssh",
			want:    "ssh://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert https with user into ssh",
			rawURL:  "https://user@host.xz/path/to/repo.git",
			scheme:  "ssh",
			want:    "ssh://user@host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:   "convert https into ssh with specified user",
			rawURL: "https://host.xz/path/to/repo.git",
			scheme: "ssh",
			opts: &Options{
				User: takePointer("user"),
			},
			want:    "ssh://user@host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:   "convert https into SCP-like URL",
			rawURL: "https://host.xz/path/to/repo.git",
			scheme: "ssh",
			opts: &Options{
				ScpLike: true,
			},
			want:    "host.xz:path/to/repo.git",
			wantErr: false,
		},
		{
			name:   "convert https into SCP-like URL with specified user",
			rawURL: "https://host.xz/path/to/repo.git",
			scheme: "ssh",
			opts: &Options{
				ScpLike: true,
				User:    takePointer("user"),
			},
			want:    "user@host.xz:path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert https into http",
			rawURL:  "https://host.xz/path/to/repo.git",
			scheme:  "http",
			want:    "http://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert https into git",
			rawURL:  "https://host.xz/path/to/repo.git",
			scheme:  "git",
			want:    "git://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert https into file",
			rawURL:  "https://host.xz/path/to/repo.git",
			scheme:  "file",
			want:    "file:///path/to/repo.git",
			wantErr: false,
		},

		// The cases for converting SCP-like syntax.
		{
			name:    "convert SCP-like URL into https",
			rawURL:  "host.xz:path/to/repo.git",
			scheme:  "https",
			want:    "https://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert SCP-like URL with user into https",
			rawURL:  "user@host.xz:path/to/repo.git",
			scheme:  "https",
			want:    "https://user@host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:   "convert SCP-like URL into https while pruning user",
			rawURL: "user@host.xz:path/to/repo.git",
			scheme: "https",
			opts: &Options{
				User: takePointer(""),
			},
			want:    "https://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert SCP-like URL into http",
			rawURL:  "host.xz:path/to/repo.git",
			scheme:  "http",
			want:    "http://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert SCP-like URL into git",
			rawURL:  "host.xz:path/to/repo.git",
			scheme:  "git",
			want:    "git://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert SCP-like URL into file",
			rawURL:  "host.xz:path/to/repo.git",
			scheme:  "file",
			want:    "file:///path/to/repo.git",
			wantErr: false,
		},

		// The cases for converting ssh.
		{
			name:    "convert ssh into https",
			rawURL:  "ssh://host.xz/path/to/repo.git",
			scheme:  "https",
			want:    "https://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert ssh with user into https",
			rawURL:  "ssh://user@host.xz/path/to/repo.git",
			scheme:  "https",
			want:    "https://user@host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:   "convert ssh into https while pruning user",
			rawURL: "ssh://user@host.xz/path/to/repo.git",
			scheme: "https",
			opts: &Options{
				User: takePointer(""),
			},
			want:    "https://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert ssh into http",
			rawURL:  "ssh://host.xz/path/to/repo.git",
			scheme:  "http",
			want:    "http://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert ssh into git",
			rawURL:  "ssh://host.xz/path/to/repo.git",
			scheme:  "git",
			want:    "git://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert ssh into file",
			rawURL:  "ssh://host.xz/path/to/repo.git",
			scheme:  "file",
			want:    "file:///path/to/repo.git",
			wantErr: false,
		},

		// The cases for converting http.
		{
			name:    "convert http into ssh",
			rawURL:  "http://host.xz/path/to/repo.git",
			scheme:  "ssh",
			want:    "ssh://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert http with user into ssh",
			rawURL:  "http://user@host.xz/path/to/repo.git",
			scheme:  "ssh",
			want:    "ssh://user@host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:   "convert http into ssh with specified user",
			rawURL: "http://host.xz/path/to/repo.git",
			scheme: "ssh",
			opts: &Options{
				User: takePointer("user"),
			},
			want:    "ssh://user@host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:   "convert http into SCP-like URL",
			rawURL: "http://host.xz/path/to/repo.git",
			scheme: "ssh",
			opts: &Options{
				ScpLike: true,
			},
			want:    "host.xz:path/to/repo.git",
			wantErr: false,
		},
		{
			name:   "convert http into SCP-like URL with specified user",
			rawURL: "http://host.xz/path/to/repo.git",
			scheme: "ssh",
			opts: &Options{
				ScpLike: true,
				User:    takePointer("user"),
			},
			want:    "user@host.xz:path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert https into https",
			rawURL:  "http://host.xz/path/to/repo.git",
			scheme:  "https",
			want:    "https://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert http into git",
			rawURL:  "http://host.xz/path/to/repo.git",
			scheme:  "git",
			want:    "git://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert http into file",
			rawURL:  "http://host.xz/path/to/repo.git",
			scheme:  "file",
			want:    "file:///path/to/repo.git",
			wantErr: false,
		},

		// The cases for converting git.
		{
			name:    "convert git into ssh",
			rawURL:  "git://host.xz/path/to/repo.git",
			scheme:  "ssh",
			want:    "ssh://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:   "convert git into ssh with specified user",
			rawURL: "git://host.xz/path/to/repo.git",
			scheme: "ssh",
			opts: &Options{
				User: takePointer("user"),
			},
			want:    "ssh://user@host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:   "convert git into SCP-like URL",
			rawURL: "git://host.xz/path/to/repo.git",
			scheme: "ssh",
			opts: &Options{
				ScpLike: true,
			},
			want:    "host.xz:path/to/repo.git",
			wantErr: false,
		},
		{
			name:   "convert git into SCP-like URL with specified user",
			rawURL: "git://host.xz/path/to/repo.git",
			scheme: "ssh",
			opts: &Options{
				ScpLike: true,
				User:    takePointer("user"),
			},
			want:    "user@host.xz:path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert git into https",
			rawURL:  "git://host.xz/path/to/repo.git",
			scheme:  "https",
			want:    "https://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert git into http",
			rawURL:  "git://host.xz/path/to/repo.git",
			scheme:  "http",
			want:    "http://host.xz/path/to/repo.git",
			wantErr: false,
		},
		{
			name:    "convert git into file",
			rawURL:  "git://host.xz/path/to/repo.git",
			scheme:  "file",
			want:    "file:///path/to/repo.git",
			wantErr: false,
		},

		// The unsupported case.
		{
			name:    "unsupported scheme given",
			rawURL:  "https://host.xz/path/to/repo.git",
			scheme:  "foo",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Convert(tt.rawURL, tt.scheme, tt.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Convert() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func takePointer(s string) *string { return &s }
