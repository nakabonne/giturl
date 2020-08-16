# giturl
`giturl` lets you convert a Git URL into the scheme you like.

## Installation

With Go
```
go get github.com/nakabonne/giturl
```

## Usage

```bash
$ giturl -h
A converter for Git URLs

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

