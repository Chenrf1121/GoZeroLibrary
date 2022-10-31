package logic

import (
	"context"
	"fmt"
	"library/common/errorx"

	"library/app/search/rpc/internal/svc"
	"library/app/search/rpc/search"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReturnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReturnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReturnLogic {
	return &ReturnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ReturnLogic) Return(in *search.ReturnReq) (*search.ReturnResp, error) {
	// todo: add your logic here and delete this line
	oneinfor, err := l.svcCtx.BooksModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return &search.ReturnResp{Status: false}, errorx.NewCodeError(10012, "更新书本数据库错误")
	}
	fmt.Println("oneinfor.CountNow 1  ===   ", oneinfor.CountNow)
	oneinfor.CountNow++
	fmt.Println("oneinfor.CountNow  2 ===   ", oneinfor.CountNow)
	err = l.svcCtx.BooksModel.Update(l.ctx, oneinfor)
	if err != nil {
		return &search.ReturnResp{Status: false}, errorx.NewCodeError(10012, "更新书本数据库错误")
	}
	return &search.ReturnResp{Status: true}, nil
}
