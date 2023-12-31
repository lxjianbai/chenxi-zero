// Code generated by goctl. DO NOT EDIT.

package gamemodel

import (
	"context"
	"strings"
	"time"
	"chenxi/pkg/dao/mysql"
	"chenxi/pkg/model"

	"gorm.io/gorm"
)

type (
	sysRoleModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *SysRole) error

		FindOne(ctx context.Context, id uint64) (*SysRole, error)
		Update(ctx context.Context, tx *gorm.DB, data *SysRole) error

		Delete(ctx context.Context, tx *gorm.DB, id uint64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
		Sharding(Sharding model.ISharding) *defaultSysRoleModel
		Builder(tx *gorm.DB) *gorm.DB
	}

	defaultSysRoleModel struct {
		conn     *gorm.DB
		table    string
		sharding model.ISharding
	}

	SysRole struct {
		Id        uint64         `gorm:"column:id"`         // 编号
		Name      string         `gorm:"column:name"`       // 名称
		UniqueKey string         `gorm:"column:unique_key"` // 唯一标识
		Remark    string         `gorm:"column:remark"`     // 备注
		MenuIds   string         `gorm:"column:menu_ids"`   // 权限集
		Status    uint64         `gorm:"column:status"`     // 0=禁用 1=开启
		CreatedAt time.Time      `gorm:"column:created_at"` // 创建时间
		UpdatedAt time.Time      `gorm:"column:updated_at"` // 更新时间
		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
	}
)

func (SysRole) TableName() string {
	return "`sys_role`"
}

func newSysRoleModel(conn *gorm.DB) *defaultSysRoleModel {
	return &defaultSysRoleModel{
		conn:  conn,
		table: "`sys_role`",
	}
}

func (m *defaultSysRoleModel) Sharding(sharding model.ISharding) *defaultSysRoleModel {
	m.sharding = sharding
	return m
}

func (m *defaultSysRoleModel) Builder(tx *gorm.DB) *gorm.DB {
	return m.scopes(tx).Model(&SysRole{})
}

func (m *defaultSysRoleModel) scopes(tx *gorm.DB) *gorm.DB {
	var db = m.conn
	if tx != nil {
		db = tx
	}
	if m.sharding == nil {
		return db
	}
	return db.Scopes(func(d *gorm.DB) *gorm.DB {
		return d.Table(strings.Trim(m.table, "`") + "_" + m.sharding.GetTableSuffix())
	})
}

func (m *defaultSysRoleModel) Insert(ctx context.Context, tx *gorm.DB, data *SysRole) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(&data).Error
	return err
}

func (m *defaultSysRoleModel) FindOne(ctx context.Context, id uint64) (*SysRole, error) {
	var resp SysRole
	err := m.scopes(nil).WithContext(ctx).Model(&SysRole{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case mysql.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysRoleModel) Update(ctx context.Context, tx *gorm.DB, data *SysRole) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(data).Error
	return err
}

func (m *defaultSysRoleModel) Delete(ctx context.Context, tx *gorm.DB, id uint64) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Delete(&SysRole{}, id).Error

	return err
}

func (m *defaultSysRoleModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.scopes(nil).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var db = m.scopes(tx)
		return fn(db)
	})
}
