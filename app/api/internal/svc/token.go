package svc

import (
	"context"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

// GetJwtToken 从ServiceContext结构体中获取JWT令牌
func (l *ServiceContext) GetJwtToken(userId uint64) (string, error) {
	// 创建JWT声明
	claims := make(jwt.MapClaims)
	// 获取当前时间
	var now = time.Now()
	// 设置令牌过期时间
	claims["exp"] = now.Unix() + l.Config.Auth.AccessExpire
	// 设置令牌签发时间
	claims["iat"] = now.Unix()
	// 设置用户ID
	claims["user_id"] = userId
	// 创建JWT令牌
	token := jwt.New(jwt.SigningMethodHS256)
	// 设置令牌声明
	token.Claims = claims
	// 用访问密钥对令牌进行签名，并返回签名后的令牌字符串
	return token.SignedString([]byte(l.Config.Auth.AccessSecret))
}

// GetUidFromCtx 从上下文中获取用户ID
func (l *ServiceContext) GetUidFromCtx(ctx context.Context) uint64 {
	var uid uint64
	if jsonUid, ok := ctx.Value("user_id").(json.Number); ok { // 从上下文中获取"user_id"对应的值，并判断是否为json.Number类型
		if int64Uid, err := jsonUid.Int64(); err == nil { // 将json.Number类型转换为int64类型
			uid = uint64(int64Uid) // 将int64类型转换为uint64类型
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err) // 记录错误日志
		}
	}
	return uid // 返回获取到的用户ID
}
