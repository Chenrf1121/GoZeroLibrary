package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"library/app/login/model"
	"library/app/login/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
