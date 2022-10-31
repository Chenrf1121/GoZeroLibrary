package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"library/app/login/rpc/loginclient"
	"library/app/search/api/internal/config"
	"library/app/search/api/internal/middleware"
	"library/app/search/model"
)

type ServiceContext struct {
	Config      config.Config
	SearchModel model.BooksModel
	UserRpc     loginclient.Login
	Example     rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		SearchModel: model.NewBooksModel(conn, c.CacheRedis),
		UserRpc:     loginclient.NewLogin(zrpc.MustNewClient(c.UserRpc)),
		Example:     middleware.NewExampleMiddleware().Handle,
	}
}
