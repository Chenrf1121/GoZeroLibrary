package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"library/app/login/rpc/login"
	"library/app/search/api/internal/svc"
	"library/app/search/api/internal/types"
	"library/app/search/model"
	"library/common/errorx"
	"strconv"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (*types.SearchResp, error) {
	// todo: add your logic here and delete this line
	userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId"))) //学号

	logx.Infof("userId: %s", userIdNumber)
	userId, err := userIdNumber.Int64()
	if err != nil {
		return nil, errorx.NewCodeError(10003, "请重新登录!")
	}

	//使用user grpc
	info, err := l.svcCtx.UserRpc.Login(l.ctx, &login.IdReq{
		Id: strconv.Itoa(int(userId)),
	})
	fmt.Println("info === ", info)
	if err != nil {
		return nil, err
	}

	logx.Infof("userId: %v", l.ctx.Value("userId"))
	booksInfo, err := l.svcCtx.SearchModel.FindOneByName(l.ctx, req.Name)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("图书馆没有这本书。")
	default:
		return nil, err
	}

	return &types.SearchResp{
		Name:      booksInfo.Name,
		Count:     int(booksInfo.Count),
		Count_now: int(booksInfo.CountNow),
	}, nil
}
