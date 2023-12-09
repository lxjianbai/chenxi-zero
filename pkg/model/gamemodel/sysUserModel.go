package gamemodel

import (
	"chenxi/pkg/dao/mysql"
	"context"

	"gorm.io/gorm"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		customSysUserLogicModel
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}

	customSysUserLogicModel interface {
		FindOneByAccount(ctx context.Context, field string, account string) (*SysUser, error)
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn *gorm.DB) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn),
	}
}

func (m *defaultSysUserModel) customCacheKeys(data *SysUser) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultSysUserModel) FindOneByAccount(ctx context.Context, field string, account string) (*SysUser, error) {
	var resp SysUser
	err := m.scopes(nil).WithContext(ctx).
		Model(&SysUser{}).
		Select(field).
		Where("`account` = ?", account).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case mysql.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
