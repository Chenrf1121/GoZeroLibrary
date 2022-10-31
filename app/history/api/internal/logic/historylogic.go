package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"library/app/borrow/rpc/borrow"
	"library/app/history/api/internal/svc"
	"library/app/history/api/internal/types"
)

type HistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HistoryLogic {
	return &HistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//获取一共借阅了什么书
func (l *HistoryLogic) History(req *types.Hisreq) (*types.HistoryListResp, error) {
	// todo: 查询借过什么书
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userid: %s", userIdNumber)
	result, _ := l.svcCtx.BorrowRpc.SearchAll(l.ctx, &borrow.BorrwoIdReq{Id: userIdNumber.String()})
	var tmplist []types.Hisresp
	for _, i := range result.List {
		var tmpresp types.Hisresp
		tmpresp.ReturnTime = i.ReturnTime
		tmpresp.BorrowTime = i.BorrowTime
		tmpresp.UserId = i.UserId
		tmpresp.BookId = i.BookId
		tmpresp.IsReturn = i.Isreturn
		tmplist = append(tmplist, tmpresp)
	}
	return &types.HistoryListResp{List: tmplist}, nil
}
