package server

import (
	"fmt"
	"github.com/zuiwuchang/king-update/cmd/server/configure"
	"github.com/zuiwuchang/king-update/cmd/server/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// Authorization 權限 認證 系統
type Authorization struct {
	Root map[string]string
	User map[string]string
}

var gAuthorization Authorization

func initAuthorization(cnf *configure.Configure) {
	root := make(map[string]string)
	user := make(map[string]string)

	var ok bool
	for i := 0; i < len(cnf.Users); i++ {
		node := &cnf.Users[i]
		switch node.Mode {
		case configure.UserModeRoot:
			if _, ok = root[node.Name]; !ok {
				root[node.Name] = node.Pwd
			}
		case configure.UserModeUser:
			if _, ok = user[node.Name]; !ok {
				user[node.Name] = node.Pwd
			}
		}
	}

	gAuthorization.Root = root
	gAuthorization.User = user
}

// AuthFunc 驗證 權限
func (a *Authorization) AuthFunc(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "unknow token")
	}
	//獲取 用戶名 密碼
	var name, pwd string
	var strs []string
	if strs, ok = md["Name"]; ok && len(strs) >= 0 {
		name = strs[0]
	} else {
		return nil, grpc.Errorf(codes.Unauthenticated, "unknow token")
	}
	if strs, ok = md["Pwd"]; ok && len(strs) >= 0 {
		pwd = strs[0]
	} else {
		return nil, grpc.Errorf(codes.Unauthenticated, "unknow token")
	}
	var power, str string
	//驗證 root
	if str, ok = gAuthorization.Root[name]; ok && str == pwd {
		power = fmt.Sprint(configure.UserModeRoot)
	} else if str, ok = gAuthorization.User[name]; ok && str == pwd {
		power = fmt.Sprint(configure.UserModeUser)
	} else {
		return nil, grpc.Errorf(codes.Unauthenticated, "unknow token")
	}

	if log.Trace != nil {
		log.Trace.Println("Authorization :", name, power)
	}

	//使用 用戶 權限
	ctx = metadata.NewIncomingContext(ctx,
		metadata.Pairs("Mode", power),
	)
	return handler(ctx, req)
}
