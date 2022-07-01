package biz

import (
	"context"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUseCase)

// {{cookiecutter.project_name | title}}Repo is a {{cookiecutter.project_name | title}} repo.
type {{cookiecutter.project_name | title}}Repo interface {
	Save(context.Context, *{{cookiecutter.project_name | title}}) (*{{cookiecutter.project_name | title}}, error)
	Update(context.Context, *{{cookiecutter.project_name | title}}) (*{{cookiecutter.project_name | title}}, error)
	FindByID(context.Context, int64) (*{{cookiecutter.project_name | title}}, error)
	ListByHello(context.Context, string) ([]*{{cookiecutter.project_name | title}}, error)
	ListAll(context.Context) ([]*{{cookiecutter.project_name | title}}, error)
}
