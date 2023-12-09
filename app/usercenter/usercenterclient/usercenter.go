// Code generated by goctl. DO NOT EDIT.
// Source: usercenter.proto

package usercenterclient

import (
	"context"

	"chenxi/app/usercenter/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	UserInfoReply   = usercenter.UserInfoReply
	UserInfoRequest = usercenter.UserInfoRequest

	Usercenter interface {
		GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

func (m *defaultUsercenter) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	client := usercenter.NewUsercenterClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}