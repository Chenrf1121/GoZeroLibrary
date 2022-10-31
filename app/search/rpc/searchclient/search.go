// Code generated by goctl. DO NOT EDIT!
// Source: search.proto

package searchclient

import (
	"context"

	"library/app/search/rpc/search"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BorrowReq  = search.BorrowReq
	BorrwoResp = search.BorrwoResp
	ReturnReq  = search.ReturnReq
	ReturnResp = search.ReturnResp
	SearchReq  = search.SearchReq
	SearchResp = search.SearchResp

	Search interface {
		Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchResp, error)
		Borrow(ctx context.Context, in *BorrowReq, opts ...grpc.CallOption) (*BorrwoResp, error)
		Return(ctx context.Context, in *ReturnReq, opts ...grpc.CallOption) (*ReturnResp, error)
	}

	defaultSearch struct {
		cli zrpc.Client
	}
)

func NewSearch(cli zrpc.Client) Search {
	return &defaultSearch{
		cli: cli,
	}
}

func (m *defaultSearch) Search(ctx context.Context, in *SearchReq, opts ...grpc.CallOption) (*SearchResp, error) {
	client := search.NewSearchClient(m.cli.Conn())
	return client.Search(ctx, in, opts...)
}

func (m *defaultSearch) Borrow(ctx context.Context, in *BorrowReq, opts ...grpc.CallOption) (*BorrwoResp, error) {
	client := search.NewSearchClient(m.cli.Conn())
	return client.Borrow(ctx, in, opts...)
}

func (m *defaultSearch) Return(ctx context.Context, in *ReturnReq, opts ...grpc.CallOption) (*ReturnResp, error) {
	client := search.NewSearchClient(m.cli.Conn())
	return client.Return(ctx, in, opts...)
}
