package main

import (
	"github.com/zuiwuchang/king-update/cmd"
	"os"
)

func main() {
	e := cmd.Execute()
	if e != nil {
		os.Exit(1)
	}
}
