package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"library/app/borrow/rpc/borrow"
	"library/app/borrow/rpc/searchall"
	"library/app/history/api/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	BorrowRpc borrow.SearchAllClient
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:    c,
		BorrowRpc: searchall.NewSearchAll(zrpc.MustNewClient(c.BorrowRpc)),
	}
}
