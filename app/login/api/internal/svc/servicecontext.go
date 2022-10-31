package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"library/app/login/api/internal/config"
	"library/app/login/model"
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
