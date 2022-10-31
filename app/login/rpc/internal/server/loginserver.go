// Code generated by goctl. DO NOT EDIT!
// Source: login-rpc.proto

package server

import (
	"context"

	"library/app/login/rpc/internal/logic"
	"library/app/login/rpc/internal/svc"
	"library/app/login/rpc/login"
)

type LoginServer struct {
	svcCtx *svc.ServiceContext
	login.UnimplementedLoginServer
}

func NewLoginServer(svcCtx *svc.ServiceContext) *LoginServer {
	return &LoginServer{
		svcCtx: svcCtx,
	}
}

func (s *LoginServer) Login(ctx context.Context, in *login.IdReq) (*login.IdResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}
