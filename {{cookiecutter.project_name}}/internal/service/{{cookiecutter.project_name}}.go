package service

import (
	"context"
	
	"{{cookiecutter.module_name}}/internal/biz"
	v1 "{{cookiecutter.module_name}}/proto/api/{{cookiecutter.project_name}}/v1"
)

// {{cookiecutter.project_name | title}}Service is a {{cookiecutter.project_name}} service.
type {{cookiecutter.project_name | title}}Service struct {
	v1.Unimplemented{{cookiecutter.project_name | title}}Server

	uc *biz.{{cookiecutter.project_name | title}}UseCase
}

// New{{cookiecutter.project_name | title}}Service new a {{cookiecutter.project_name}} service.
func New{{cookiecutter.project_name | title}}Service(uc *biz.{{cookiecutter.project_name | title}}UseCase) *{{cookiecutter.project_name | title}}Service {
	return &{{cookiecutter.project_name | title}}Service{uc: uc}
}

// SayHello implements helloworld.{{cookiecutter.project_name | title}}Server.
func (s *{{cookiecutter.project_name | title}}Service) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	g, err := s.uc.Create{{cookiecutter.project_name | title}}(ctx, &biz.{{cookiecutter.project_name | title}}{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
}
