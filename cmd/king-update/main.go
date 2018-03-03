package main

import (
	"flag"
	"fmt"
)

func main() {
	var h, init, create bool

	flag.BoolVar(&h, "h", false, "show help")
	flag.BoolVar(&init, "init", false, fmt.Sprintf("init package file : %s %s", PackageFile, RuleFile))
	flag.BoolVar(&create, "create", false, fmt.Sprintf("create update file : %s", UpdateFile))

	flag.Parse()

	if init {
		Init()
		return
	} else if create {
		Create()
		return
	}
	flag.PrintDefaults()
}
