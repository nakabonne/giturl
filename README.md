# giturl
A tiny parser for [Git URLs](https://git-scm.com/docs/git-clone#_git_urls).

In addition to ssh, git, http, and https protocols, it also supports SCP-like URLs.

```go
package main

import (
	"fmt"

	"github.com/nakabonne/giturl"
)

func main() {
	u := giturl.Parse("git@github.com:org/repo.git")
	fmt.Printf("%#v", u)
	/*
		&url.URL{
			Scheme: "ssh",
			User: &url.Userinfo{
				username:    "git",
			},
			Host:       "github.com",
			Path:       "/org/repo.git",
		}
	*/
}
```
