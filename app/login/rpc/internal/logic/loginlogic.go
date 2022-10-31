package logic

import (
	"context"
	"strconv"

	"library/app/login/rpc/internal/svc"
	"library/app/login/rpc/login"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *login.IdReq) (*login.IdResp, error) {
	// todo: add your logic here and delete this line
	//	id, err := strconv.ParseInt(in.Id, 10, 64)
	one, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &login.IdResp{Id: strconv.Itoa(int(one.Id)), Number: one.Number,
		Name: one.Name, Gender: one.Gender}, nil
}
