package gamemodel

import "gorm.io/gorm"

type GameModel struct {
	SysUserModel SysUserModel
	SysRoleModel SysRoleModel
	SysMenuModel SysMenuModel
}

func NewGameModel(db *gorm.DB) *GameModel {
	return &GameModel{
		SysUserModel: NewSysUserModel(db),
		SysRoleModel: NewSysRoleModel(db),
		SysMenuModel: NewSysMenuModel(db),
	}
}
