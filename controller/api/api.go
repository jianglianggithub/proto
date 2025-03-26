package api

import (
	"context"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/jianglianggithub/proto/api/order"
)

type Controller struct {
	order.UnimplementedOrderInfoServer
}

func Register(s *grpcx.GrpcServer) {
	order.RegisterOrderInfoServer(s.Server, &Controller{})
}

func (*Controller) Create(ctx context.Context, req *order.CreateReq) (res *order.CreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Create1(ctx context.Context, req *order.CreateReq) (res *order.CreateRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
