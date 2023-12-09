package user

import (
	"context"

	"chenxi/app/api/internal/svc"
	"chenxi/app/api/internal/types"
	"chenxi/app/usercenter/usercenter"
	"chenxi/pkg/utils/xerror"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.EmptyRequest) (resp *types.UserInfoResponse, err error) {
	var uid = l.svcCtx.GetUidFromCtx(l.ctx)
	if uid == 0 {
		return nil, xerror.NewCodeError(xerror.AuthCode, "用户异常")
	}
	_, err = l.svcCtx.UserCenterRpc.GetUserInfo(l.ctx, &usercenter.UserInfoRequest{UserId: uid})
	if err != nil {
		return nil, err
	}
	return
}
