# ScummAtlas

## How to run this
I know the project doesn't follow the standard Golang structure. It's on the todo list :)
```bash
git clone git@github.com:ktzar/scummatlas.git
cd scummatlas
export GO111MODULE="auto"
export GOPATH=`pwd`
go run src/scummatlas/main/scummatlas.go -gamedir=/path/to/games/monkey2 -outputdir out
```

<img src="https://api.travis-ci.org/ktzar/scummatlas.svg?branch=master"/>

Scumm games parser and HTML Atlas generator written in Golang.

The aim of this software is to provide an easy to understand, testable, and organised software that unpacks games using the SCUMM engine.

It creates a bunch of HTML files that software that aim to show the inner aspects of how these beloved games were implemented.

Build status: https://travis-ci.org/ktzar/scummatlas

TODO
- Improve parsing local and object scripts
- ~~Parse costumes~~
- ~~Have a game detector~~
- Add ordering capabilities to table view
- Use c64 font in some places like titles
