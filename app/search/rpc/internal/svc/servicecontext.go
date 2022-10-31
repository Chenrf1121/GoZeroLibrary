package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"library/app/search/rpc/internal/config"
)
import "library/app/search/model"

type ServiceContext struct {
	Config     config.Config
	BooksModel model.BooksModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		BooksModel: model.NewBooksModel(conn, c.CacheRedis),
	}
}
