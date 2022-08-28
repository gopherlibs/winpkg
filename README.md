# winpkg [![CI Status](https://circleci.com/gh/gopherlibs/winpkg.svg?style=shield)](https://app.circleci.com/pipelines/github/gopherlibs/winpkg) [![Go Report Card](https://goreportcard.com/badge/github.com/gopherlibs/winpkg)](https://goreportcard.com/report/github.com/gopherlibs/winpkg) [![GoDoc](https://godoc.org/github.com/gopherlibs/appindicator?status.svg)](https://godoc.org/github.com/gopherlibs/appindicator) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/gopherlibs/winpkg/trunk/LICENSE)

`winpkg` is a Go (Golang) module to pull the latest version of a Chocolatey package.
Chocolatey here is a Windows package manager.


## Requirements

This Go module is intended to support the ~~two~~ most recent ~~minor releases~~ release of Go.
Currently this is Go v1.18.x and v1.19.x.


## Installation

`winpkg` is a Go module intended to be used as a library.
So the best way to "install" it is to import into your own code and then run `go mod tidy` to get the source code downloaded.

```go
import(
	"github.com/gopherlibs/winpkg/winpkg"
)
```


## Usage


```go
package main

import (
	"fmt"

	"github.com/gopherlibs/winpkg/winpkg"
)

func main() {

	// Let's get the latest version of the 'git' package
	if ver, err := winpkg.GetVersion( "git" ); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ver)
	}
}
```


## Development

This module is written and tested with Go v1.18+ in mind.
`go fmt` is your friend.
Please feel free to open Issues and PRs are you see fit.
Any PR that requires a good amount of work or is a significant change, it would be best to open an Issue to discuss the change first.


## License & Credits

This module was written by Ricardo N Feliciano (FelicianoTech).
This repository is licensed under the MIT license.
This repo's license can be found [here](./LICENSE).
