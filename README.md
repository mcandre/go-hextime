# go-hextime - a Go port of hextime

# EXAMPLE

```console
$ hextime
e_40_9
```

# DOWNLOAD

https://github.com/mcandre/go-hextime/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/go-hextime

# RUNTIME REQUIREMENTS

(None)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.9+
* [Docker](https://www.docker.com/)
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [goxcart](https://github.com/mcandre/goxcart) (e.g., `github.com/mcandre/goxcart/...`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)

## Recommended

* [karp](https://github.com/mcandre/karp) (e.g., `go get github.com/mcandre/karp/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```console
$ go get github.com/mcandre/go-hextime/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```console
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone https://github.com/mcandre/go-hextime.git $GOPATH/src/github.com/mcandre/go-hextime
$ cd "$GOPATH/src/github.com/mcandre/go-hextime"
$ git submodule update --init --recursive
$ go install ./...
```

# TEST

```console
$ mage test
```

# PORT

```console
$ mage port
```

# LINT

Keep the code tidy:

```console
$ mage lint
```
