package logic

import (
	"context"
	"errors"
	"library/app/login/model"
	"strconv"
	"time"

	"library/app/login/api/internal/svc"
	"library/app/login/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Sing_upLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSing_upLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Sing_upLogic {
	return &Sing_upLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Sing_upLogic) Sing_up(req *types.SignupReq) (*types.SingupResq, error) {
	// todo: add your logic here and delete this line
	//	id,_ := strconv.Atoi(req.Id)
	number, _ := strconv.ParseInt(req.Id, 10, 64)
	_, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, strconv.Itoa(int(number)))
	if err == nil {
		return &types.SingupResq{
			Ok: false,
		}, errors.New("账号已经存在，请登录。")
	}

	data := model.User{
		Number:     req.Id,
		Name:       req.Name,
		Password:   req.Password,
		Gender:     req.Gender,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, &data)
	if err != nil {
		return nil, err
	}
	return &types.SingupResq{
		Ok: true,
	}, nil
}
