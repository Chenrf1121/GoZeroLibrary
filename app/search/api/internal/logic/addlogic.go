package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"library/app/login/rpc/login"
	"library/app/search/api/internal/svc"
	"library/app/search/api/internal/types"
	"library/app/search/model"
	"library/common/errorx"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.AddReq) (*types.AddResp, error) {
	// todo: add your logic here and delete this line

	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
	logx.Infof("userId is :%s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(10003, "请重新登录!")
	}
	info, err := l.svcCtx.UserRpc.Login(l.ctx, &login.IdReq{
		Id: strconv.Itoa(int(userId)),
	})
	logx.Infof("user name = %+V", info)
	if err != nil {
		return nil, errorx.NewCodeError(10004, "远程调用rpc错误!")
	}

	booksInfo, err := l.svcCtx.SearchModel.FindOneByName(l.ctx, req.Name)
	if err == nil {
		booksInfo.Count += int64(req.Count)
		booksInfo.CountNow += int64(req.Count)
		err1 := l.svcCtx.SearchModel.Update(l.ctx, booksInfo)
		if err1 != nil {
			return nil, err1
		} else {
			return &types.AddResp{
				Ok: true,
			}, nil
		}
	} else if err == model.ErrNotFound {
		tmp_bookInfo := model.Books{Count: int64(req.Count), CountNow: int64(req.Count), Name: req.Name}
		_, err1 := l.svcCtx.SearchModel.Insert(l.ctx, &tmp_bookInfo)
		if err1 != nil {
			return nil, err1
		}
	} else {
		return nil, err
	}
	return &types.AddResp{
		Ok: true,
	}, nil
}
