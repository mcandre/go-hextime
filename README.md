# go-hextime - a Go port of hextime

# EXAMPLE

```
$ hextime
e_40_9
```

# DOWNLOAD

https://github.com/mcandre/go-hextime/releases

# DEVELOPMENT REQUIREMENTS

* [Go](https://golang.org) 1.7+ with [$GOPATH configured](https://gist.github.com/mcandre/ef73fb77a825bd153b7836ddbd9a6ddc)

## Optional

* [coreutils](https://www.gnu.org/software/coreutils/coreutils.html)
* [make](https://www.gnu.org/software/make/)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [gox](https://github.com/mitchellh/gox) (e.g. `go get github.com/mitchellh/gox`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/go-hextime/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone git@github.com:mcandre/go-hextime.git $GOPATH/src/github.com/mcandre/go-hextime
$ sh -c "cd $GOPATH/src/github.com/mcandre/go-hextime/cmd/hextime && go install"
```

# PORT

```
$ make port
```

# LINT

Keep the code tidy:

```
$ make lint
```

# GIT HOOKS

See `hooks/`.
