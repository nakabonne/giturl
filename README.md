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
```

## Examples

```bash
$ giturl https git@github.com:org/repo.git
https://github.com/org/repo.git

$ giturl ssh https://github.com/org/repo.git
ssh://github.com/org/repo.git

$ giturl ssh --scp-like --user=git https://github.com/org/repo.git
git@github.com:org/repo.git
```

