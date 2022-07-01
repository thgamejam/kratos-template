package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"{{cookiecutter.module_name}}/internal/conf"
	"{{cookiecutter.module_name}}/internal/service"
	v1 "{{cookiecutter.module_name}}/proto/api/{{cookiecutter.project_name}}/v1"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, service *service.{{cookiecutter.project_name | title}}Service, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.Register{{cookiecutter.project_name | title}}Server(srv, service)
	return srv
}
