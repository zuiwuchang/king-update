package cmd

import (
	"github.com/zuiwuchang/king-update/cmd/server"
	"github.com/zuiwuchang/king-update/cmd/tools"
)

func init() {
	cmd := initRoot()
	tools.InitSubCommand(cmd)
	server.InitSubCommand(cmd)
}
