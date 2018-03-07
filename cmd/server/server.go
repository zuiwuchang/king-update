package server

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zuiwuchang/king-update/cmd/server/configure"
	"github.com/zuiwuchang/king-update/cmd/server/log"
	pb "github.com/zuiwuchang/king-update/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"strings"
)

const (
	// ConfigureFile 配置 檔案
	ConfigureFile = "server.jsonnet"
)

func initServer() (cmd *cobra.Command) {
	var addr string
	var cnfFile string
	cmd = &cobra.Command{
		Use:   "server",
		Short: "run server daemon",
		Run: func(cmd *cobra.Command, args []string) {
			//初始化 配置
			var e = configure.Init(cnfFile)
			if e != nil {
				fmt.Println(e)
				os.Exit(1)
			}
			cnf := configure.GetConfigure()

			//命令 參數
			addr = strings.TrimSpace(addr)
			if addr != "" {
				cnf.Addr = addr
			}
			//運行 main
			serverMain(cnf)
		},
	}
	flags := cmd.Flags()
	flags.StringVarP(
		&addr,
		"addr",
		"a",
		"",
		"tcp bind address",
	)
	flags.StringVarP(
		&cnfFile,
		"configure",
		"c",
		ConfigureFile,
		"server configure file path",
	)
	return
}
func serverMain(cnf *configure.Configure) {
	//初始化 日誌
	log.Init()

	//運行 服務器
	runAsGRPC(cnf)
}
func runAsGRPC(cnf *configure.Configure) {
	opts := make([]grpc.ServerOption, 0, 2)
	// 驗證 token
	initAuthorization(cnf)
	opts = append(opts, grpc.UnaryInterceptor(gAuthorization.AuthFunc))

	// 是否 使用 tls
	var h2 bool
	if cnf.Crt != "" && cnf.Key != "" {
		h2 = true
		creds, e := credentials.NewServerTLSFromFile(cnf.Crt, cnf.Key)
		if e != nil {
			if log.Fault != nil {
				log.Fault.Println(e)
			}
			os.Exit(1)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	//創建 監聽 Listener
	l, e := net.Listen("tcp", cnf.Addr)
	if e != nil {
		if log.Fault != nil {
			log.Fault.Println()
		}
		os.Exit(1)
	}
	if log.Info != nil {
		if h2 {
			log.Info.Println("h2 work at", cnf.Addr)
		} else {
			log.Info.Println("h2c work at", cnf.Addr)
		}
	}

	//創建 rpc 服務器
	s := grpc.NewServer(opts...)

	//註冊 服務
	pb.RegisterServiceBasicServer(s, &ServiceBasic{})

	//註冊 反射 到 服務 路由
	reflection.Register(s)

	//讓 rpc 在 Listener 上 工作
	if e := s.Serve(l); e != nil {
		if log.Fault != nil {
			log.Fault.Println()
		}
		os.Exit(1)
	}
}
