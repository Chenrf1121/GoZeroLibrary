package logic

import (
	"context"
	"library/app/search/rpc/internal/svc"
	"library/app/search/rpc/search"
	"library/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchLogic) Search(in *search.SearchReq) (*search.SearchResp, error) {
	// todo: add your logic here and delete this line
	one, err := l.svcCtx.BooksModel.FindOneByName(l.ctx, in.Name)
	if err != nil {
		return nil, errorx.NewCodeError(10004, "图书馆没有这本书!")
	}
	return &search.SearchResp{
		Id:       one.Id,
		CountNow: int32(one.CountNow),
	}, nil
}
