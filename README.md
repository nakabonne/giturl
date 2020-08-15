# giturl
A converter for Git URLs.

## Installation

With Go
```
go get github.com/nakabonne/giturl/cmd/giturl
```

## Usage

```bash
$ giturl -h
usage: giturl [<flag> ...] <Git URL>
      --no-user         prune user from the given URL
  -s, --scheme string   convert to the given schema: ssh|http|https|git|file (default "ssh")
      --scp-like        emit scp-like syntax (available only when --schema=ssh)
  -u, --user string     set user
```

## Example

```bash
$ giturl --scheme=ssh https://github.com/org/repo.git
ssh://github.com/org/repo.git

$ giturl --scheme=ssh --scp-like --user=git https://github.com/org/repo.git
git@github.com:org/repo.git

$ giturl --scheme=https --no-user git@github.com:org/repo.git
https://github.com/org/repo.git

$ giturl --scheme=git https://github.com/org/repo.git
git://github.com/org/repo.git
```

## Supported protocols
- ssh
- git
- http
- https
- file
