package main

import (
	"fmt"
	"os"

	"github.com/gopherlibs/winpkg/winpkg"
)

func main() {
	if ver, err := winpkg.GetVersion(os.Args[1]); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ver)
	}
}
