package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	//AppName 程式 名稱
	AppName = "king-update"
	// Version 當前 版本
	Version = "v0.0.1"
)

var rootCommand *cobra.Command

func initRoot() (cmd *cobra.Command) {
	var flagV bool
	cmd = &cobra.Command{
		Use:   AppName,
		Short: "Automatically release software to customers (GPL-3.0)",
		Long: `Automatically release software to customers (GPL-3.0)
	
   written by king
      github    -- https://github.com/zuiwuchang/king-update
   king's blog  -- http://blog.king011.com/
   king's email -- zuiwuchang@gmail.com`,
		Run: func(cmd *cobra.Command, args []string) {
			if flagV {
				fmt.Println(AppName, Version)
			}
		},
	}
	flags := cmd.Flags()
	flags.BoolVarP(&flagV,
		"version",
		"v",
		false,
		"show version",
	)

	rootCommand = cmd
	return
}

// Execute 執行 命令行 程式
func Execute() error {
	return rootCommand.Execute()
}
