package logic

import (
	"context"

	"chenxi/app/usercenter/internal/svc"
	"chenxi/app/usercenter/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *usercenter.UserInfoRequest) (*usercenter.UserInfoReply, error) {
	// todo: add your logic here and delete this line

	return &usercenter.UserInfoReply{}, nil
}
