# giturl

[![codecov](https://codecov.io/gh/nakabonne/giturl/branch/master/graph/badge.svg)](https://codecov.io/gh/nakabonne/giturl)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/nakabonne/giturl/pkg/converter)

`giturl` lets you convert [Git URLs](https://git-scm.com/docs/git-clone#_git_urls) into the scheme you like.

## Installation

You can download a binary release [here](https://github.com/nakabonne/giturl/releases).

With Homebrew
```
brew install nakabonne/giturl/giturl
```

With Go
```
go get github.com/nakabonne/giturl
```

With Docker
```
docker run --rm nakabonne/giturl giturl
```

## Examples

```bash
# Conversion of https to ssh
$ giturl ssh https://github.com/org/repo.git
ssh://github.com/org/repo.git

# Conversion of SCP-like ssh syntax to https
$ giturl https --no-user git@github.com:org/repo.git
https://github.com/org/repo.git

# Conversion of https to SCP-like ssh syntax
$ giturl ssh --scp-like --user=git https://github.com/org/repo.git
git@github.com:org/repo.git
```

### Available commands

```
$ giturl -h
Converts Git URLs into the scheme you like.

Usage:
  giturl [command]

Available Commands:
  file        Convert into file syntax
  git         Convert into git syntax
  help        Help about any command
  http        Convert into http syntax
  https       Convert into https syntax
  ssh         Convert into ssh syntax
  version     Print the current version
```

## Supported schemes
SSH, Git, HTTP, and HTTPS protocols are available as Git URLs. The SSH protocol also supports an alternative SCP-like syntax:

- `ssh://[user@]host.xz[:port]/path/to/repo.git/`
- `[user@]host.xz:path/to/repo.git/`
- `git://host.xz[:port]/path/to/repo.git/`
- `http[s]://host.xz[:port]/path/to/repo.git/`
- `file:///path/to/repo.git/` (for local repositories)

See more: https://git-scm.com/docs/git-clone#_git_urls

## License
[MIT][license] © [Ryo Nakao][website]

[license]: /LICENSE
[website]: https://nakabonne.dev
