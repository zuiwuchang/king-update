package server

import (
	pb "github.com/zuiwuchang/king-update/protoc"
	"golang.org/x/net/context"
)

// ServiceBasic ...
type ServiceBasic struct {
}

// Ping 實現 服務器 接口
func (*ServiceBasic) Ping(ctx context.Context, in *pb.BasicPingRequest) (*pb.BasicPingReply, error) {
	return &pb.BasicPingReply{}, nil
}
