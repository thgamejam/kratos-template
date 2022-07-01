package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

var (
// ErrUserNotFound is user not found.
// ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// {{cookiecutter.project_name | title}} is a {{cookiecutter.project_name | title}} model.
type {{cookiecutter.project_name | title}} struct {
	Hello string
}

// {{cookiecutter.project_name | title}}UseCase is a {{cookiecutter.project_name | title}} use case.
type {{cookiecutter.project_name | title}}UseCase struct {
	repo {{cookiecutter.project_name | title}}Repo
	log  *log.Helper
}

// New{{cookiecutter.project_name | title}}UseCase new a {{cookiecutter.project_name | title}} use case.
func New{{cookiecutter.project_name | title}}UseCase(repo {{cookiecutter.project_name | title}}Repo, logger log.Logger) *{{cookiecutter.project_name | title}}UseCase {
	return &{{cookiecutter.project_name | title}}UseCase{repo: repo, log: log.NewHelper(logger)}
}

// Create{{cookiecutter.project_name | title}} creates a {{cookiecutter.project_name | title}}, and returns the new {{cookiecutter.project_name | title}}.
func (uc *{{cookiecutter.project_name | title}}UseCase) Create{{cookiecutter.project_name | title}}(ctx context.Context, g *{{cookiecutter.project_name | title}}) (*{{cookiecutter.project_name | title}}, error) {
	uc.log.WithContext(ctx).Infof("Create{{cookiecutter.project_name | title}}: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}
