package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"{{cookiecutter.module_name}}/internal/biz"
)

type {{cookiecutter.project_name}}Repo struct {
	data *Data
	log  *log.Helper
}

// New{{cookiecutter.project_name | title}}Repo .
func New{{cookiecutter.project_name | title}}Repo(data *Data, logger log.Logger) biz.{{cookiecutter.project_name | title}}Repo {
	return &{{cookiecutter.project_name}}Repo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *{{cookiecutter.project_name}}Repo) Save(ctx context.Context, b *biz.{{cookiecutter.project_name | title}}) (*biz.{{cookiecutter.project_name | title}}, error) {
	return b, nil
}

func (r *{{cookiecutter.project_name}}Repo) Update(ctx context.Context, b *biz.{{cookiecutter.project_name | title}}) (*biz.{{cookiecutter.project_name | title}}, error) {
	return b, nil
}

func (r *{{cookiecutter.project_name}}Repo) FindByID(context.Context, int64) (*biz.{{cookiecutter.project_name | title}}, error) {
	return nil, nil
}

func (r *{{cookiecutter.project_name}}Repo) ListByHello(context.Context, string) ([]*biz.{{cookiecutter.project_name | title}}, error) {
	return nil, nil
}

func (r *{{cookiecutter.project_name}}Repo) ListAll(context.Context) ([]*biz.{{cookiecutter.project_name | title}}, error) {
	return nil, nil
}
