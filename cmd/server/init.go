package server

import (
	"github.com/spf13/cobra"
)

// InitSubCommand 初始化 子命令
func InitSubCommand(root *cobra.Command) {
	root.AddCommand(
		initServer(),
	)
}
