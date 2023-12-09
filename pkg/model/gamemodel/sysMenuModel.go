package gamemodel

import (
	"gorm.io/gorm"
)

var _ SysMenuModel = (*customSysMenuModel)(nil)

type (
	// SysMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMenuModel.
	SysMenuModel interface {
		sysMenuModel
		customSysMenuLogicModel
	}

	customSysMenuModel struct {
		*defaultSysMenuModel
	}

	customSysMenuLogicModel interface {
	}
)

// NewSysMenuModel returns a model for the database table.
func NewSysMenuModel(conn *gorm.DB) SysMenuModel {
	return &customSysMenuModel{
		defaultSysMenuModel: newSysMenuModel(conn),
	}
}

func (m *defaultSysMenuModel) customCacheKeys(data *SysMenu) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
