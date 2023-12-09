package login

import (
	"context"
	"strings"
	"time"

	"chenxi/app/cms/internal/svc"
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/cpb"
	cmscachekey "chenxi/pkg/global/cmsCacheKey"
	"chenxi/pkg/utils"
	"chenxi/pkg/utils/xerror"

	"github.com/gookit/goutil/errorx"
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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	if len(strings.TrimSpace(req.Account)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		return nil, xerror.NewDefaultError("账号或密码错误")
	}

	var field = "id,nickname,password,avatar,status"
	sysUser, err := l.svcCtx.GameModel.SysUserModel.FindOneByAccount(l.ctx, field, req.Account)
	if err != nil {
		return nil, xerror.NewDefaultError("账号或密码错误")
	}
	if sysUser.Status != int64(cpb.SysUserStatusEnum_SysUserStatusEnumNormal) {
		return nil, xerror.NewDefaultError("该账号已锁定")
	}
	if !utils.ComparePasswords(sysUser.Password, req.Password) {
		return nil, xerror.NewDefaultError("账号或密码错误")
	}
	token, err := l.svcCtx.GetJwtToken(sysUser.Id)
	if err != nil {
		// 记录token生成失败的错误日志并返回服务器异常的错误信息
		l.Logger.Error(errorx.Wrap(err, "token生成失败"))
		return nil, xerror.NewCodeError(xerror.InternalCode, "服务器异常")
	}
	_, err = l.svcCtx.Redis.SetnxEx(cmscachekey.SysUserToken(uint32(sysUser.Id)), "1", int(l.svcCtx.Config.Auth.AccessExpire))
	if err != nil {
		l.Logger.Error(errorx.Wrap(err, "token缓存失败"))
		return nil, xerror.NewCodeError(xerror.InternalCode, "服务器异常")
	}
	resp = new(types.LoginResponse)

	resp.Token = token
	resp.Name = sysUser.Name
	resp.Avatar = sysUser.Avatar
	// 异步更新用户登录信息
	go l.updateUserLoginInfo(sysUser.Id)
	return
}

// updateUserLoginInfo 方法用于更新用户的登录信息
func (l *LoginLogic) updateUserLoginInfo(userId uint64) {
	// 使用 svcCtx 中的 UsersModel 对象构建更新操作，更新用户的最后登录时间为当前时间戳
	if err := l.svcCtx.GameModel.SysUserModel.Builder(nil).Where("id = ?", userId).Update("login_at", time.Now()).Error; err != nil {
		// 如果更新操作出现错误，则记录错误日志，包括用户ID和错误信息
		l.Logger.WithFields(logx.Field("user_id", userId)).Error(errorx.Wrap(err, "更新用户登录信息失败"))
	}
}
