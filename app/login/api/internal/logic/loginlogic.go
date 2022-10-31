package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"library/app/login/model"
	"library/common/errorx"
	"strconv"
	"strings"
	"time"

	"library/app/login/api/internal/svc"
	"library/app/login/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))

}
func (l *LoginLogic) Login(req *types.LoginReq) (*types.LoginResq, error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(req.Id)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, errorx.NewCodeError(10000, "请输入用户名和密码！")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, req.Id)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewCodeError(10001, "用户名不存在，请注册！")
	default:
		return nil, err
	}
	if userInfo.Password != req.Password {
		return nil, errorx.NewCodeError(10003, "密码不正确！")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	userid, _ := strconv.Atoi(userInfo.Number)
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(userid))
	if err != nil {
		return nil, err
	}

	return &types.LoginResq{
		Id:           userInfo.Number,
		Name:         userInfo.Name,
		Gender:       userInfo.Gender,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}
