package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"library/app/borrow/api/internal/config"
	"library/app/borrow/model"
	"library/app/login/rpc/login"
	"library/app/login/rpc/loginclient"
	"library/app/search/rpc/search"
	"library/app/search/rpc/searchclient"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     login.LoginClient
	BorrowModel model.BorrowModel
	SearchRpc   search.SearchClient
	//	BooksModel  sem.BooksModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		BorrowModel: model.NewBorrowModel(conn, c.CacheRedis),
		UserRpc:     loginclient.NewLogin(zrpc.MustNewClient(c.UserRpc)),
		SearchRpc:   searchclient.NewSearch(zrpc.MustNewClient(c.SearchRpc)),
	}
}
