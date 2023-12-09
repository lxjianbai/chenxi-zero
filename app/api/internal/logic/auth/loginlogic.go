package auth

import (
	"context"
	"strings"
	"time"

	"chenxi/app/api/internal/svc"
	"chenxi/app/api/internal/types"
	"chenxi/pkg/cpb"
	"chenxi/pkg/utils"
	"chenxi/pkg/utils/xerror"

	"github.com/gookit/goutil/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
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

// Login方法用于处理用户登录逻辑
func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// 检查账号和密码是否为空
	if strings.TrimSpace(req.Account) == "" || strings.TrimSpace(req.Password) == "" {
		// 返回账号或密码错误的错误信息
		return nil, xerror.NewCodeError(xerror.AuthCode, "账号或密码错误")
	}
	// 定义需要查询的字段
	var field = "id,name,phone,password,is_delete,avatar,is_show,good_number"
	// 根据手机号查询用户信息
	user, err := l.svcCtx.CloudModel.UsersModel.FindOneByPhone(l.ctx, field, req.Account)
	if err != nil {
		// 记录用户不存在的错误日志并返回账号或密码错误的错误信息
		l.Logger.WithFields(logx.Field("account", req.Account)).Error(errorx.Wrap(err, "用户不存在"))
		return nil, xerror.NewCodeError(xerror.AuthCode, "账号或密码错误")
	}
	// 比较用户输入的密码和数据库中的密码是否一致
	if !utils.ComparePasswords(user.Password.String, req.Password) && l.svcCtx.Config.Mode == service.ProMode {
		// 记录密码错误的错误日志并返回账号或密码错误的错误信息
		l.Logger.WithFields(logx.Field("account", req.Account)).Error(errorx.Wrap(err, "密码错误"))
		return nil, xerror.NewCodeError(xerror.AuthCode, "账号或密码错误")
	}
	// 检查用户是否被删除
	if user.IsDelete == uint64(cpb.UserDeleteStatusEnum_UserDeleteStatusEnumDelete) {
		// 记录用户被删除的错误日志并返回该账户已被注销的错误信息
		l.Logger.WithFields(logx.Field("account", req.Account)).Error(errorx.Wrap(err, "用户被删除"))
		return nil, xerror.NewCodeError(xerror.AuthCode, "该账户已被注销")
	}
	// 生成JWT token
	token, err := l.svcCtx.GetJwtToken(user.Id)
	if err != nil {
		// 记录token生成失败的错误日志并返回服务器异常的错误信息
		l.Logger.Error(errorx.Wrap(err, "token生成失败"))
		return nil, xerror.NewCodeError(xerror.InternalCode, "服务器异常")
	}
	// 构造登录响应对象
	var reply = new(types.LoginResponse)
	reply.Token = token
	reply.Id = user.Id
	reply.Name = user.Name.String
	reply.Avatar = user.Avatar.String
	reply.GoodNumber = user.GoodNumber.String
	// 异步更新用户登录信息
	go l.updateUserLoginInfo(user.Id)
	// 返回登录响应和空错误
	return reply, nil
}

// updateUserLoginInfo 方法用于更新用户的登录信息
func (l *LoginLogic) updateUserLoginInfo(userId uint64) {
	// 使用 svcCtx 中的 UsersModel 对象构建更新操作，更新用户的最后登录时间为当前时间戳
	if err := l.svcCtx.CloudModel.UsersModel.Builder(nil).Where("id = ?", userId).Update("last_time", time.Now().Unix()).Error; err != nil {
		// 如果更新操作出现错误，则记录错误日志，包括用户ID和错误信息
		l.Logger.WithFields(logx.Field("user_id", userId)).Error(errorx.Wrap(err, "更新用户登录信息失败"))
	}
}
