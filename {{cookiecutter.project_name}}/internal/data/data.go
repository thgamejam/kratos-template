package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"{{cookiecutter.module_name}}/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	New{{cookiecutter.project_name | title}}Repo,
)

// Data .
type Data struct {
	// TODO 封装的数据库客户端
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		// TODO 装填数据库客户端
	}, cleanup, nil
}
