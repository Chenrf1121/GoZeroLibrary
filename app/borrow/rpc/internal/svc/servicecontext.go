package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"library/app/borrow/model"
	"library/app/borrow/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	BorrowModel model.BorrowModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		BorrowModel: model.NewBorrowModel(conn, c.CacheRedis),
	}
}
