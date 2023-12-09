package gamemodel

import (
	"gorm.io/gorm"
)

var _ SysRoleModel = (*customSysRoleModel)(nil)

type (
	// SysRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleModel.
	SysRoleModel interface {
		sysRoleModel
		customSysRoleLogicModel
	}

	customSysRoleModel struct {
		*defaultSysRoleModel
	}

	customSysRoleLogicModel interface {
	}
)

// NewSysRoleModel returns a model for the database table.
func NewSysRoleModel(conn *gorm.DB) SysRoleModel {
	return &customSysRoleModel{
		defaultSysRoleModel: newSysRoleModel(conn),
	}
}

func (m *defaultSysRoleModel) customCacheKeys(data *SysRole) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
