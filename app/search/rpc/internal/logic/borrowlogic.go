package logic

import (
	"context"
	"library/common/errorx"

	"library/app/search/rpc/internal/svc"
	"library/app/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type BorrowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBorrowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BorrowLogic {
	return &BorrowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BorrowLogic) Borrow(in *search.BorrowReq) (*search.BorrwoResp, error) {
	// todo: add your logic here and delete this line
	bookinfo, err := l.svcCtx.BooksModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return &search.BorrwoResp{Status: false}, errorx.NewCodeError(10010, "借书出错")
	}
	bookinfo.CountNow--
	err = l.svcCtx.BooksModel.Update(l.ctx, bookinfo)
	if err != nil {
		return &search.BorrwoResp{Status: false}, errorx.NewCodeError(10009, "借书错误")
	}
	return &search.BorrwoResp{Status: true}, nil
}
