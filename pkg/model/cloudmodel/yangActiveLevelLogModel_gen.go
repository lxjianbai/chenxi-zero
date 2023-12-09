// Code generated by goctl. DO NOT EDIT.

package cloudmodel

import (
	"context"
	"strings"
	"chenxi/pkg/dao/mysql"
	"chenxi/pkg/model"

	"gorm.io/gorm"
)

type (
	yangActiveLevelLogModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *YangActiveLevelLog) error

		FindOne(ctx context.Context, id uint64) (*YangActiveLevelLog, error)
		Update(ctx context.Context, tx *gorm.DB, data *YangActiveLevelLog) error

		Delete(ctx context.Context, tx *gorm.DB, id uint64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
		Sharding(Sharding model.ISharding) *defaultYangActiveLevelLogModel
		Builder(tx *gorm.DB) *gorm.DB
	}

	defaultYangActiveLevelLogModel struct {
		conn     *gorm.DB
		table    string
		sharding model.ISharding
	}

	YangActiveLevelLog struct {
		Id         uint64  `gorm:"column:id"`
		Uid        uint64  `gorm:"column:uid"`         // 用户主键
		SourceUid  uint64  `gorm:"column:source_uid"`  // 变更来源用户主键
		Event      string  `gorm:"column:event"`       // 事件标识
		Rate       float64 `gorm:"column:rate"`        // 参与计算的比例
		ResultRate float64 `gorm:"column:result_rate"` // 运算后的比例
		LinkId     string  `gorm:"column:link_id"`     // 关联主键
		Pm         int64   `gorm:"column:pm"`          // 增减类别。0：减少，1：增加
		Remark     string  `gorm:"column:remark"`      // 备注内容
		Ymd        string  `gorm:"column:ymd"`         // 年月日
		CreateTime uint64  `gorm:"column:create_time"`
		UpdateTime uint64  `gorm:"column:update_time"`
	}
)

func (YangActiveLevelLog) TableName() string {
	return "`yang_active_level_log`"
}

func newYangActiveLevelLogModel(conn *gorm.DB) *defaultYangActiveLevelLogModel {
	return &defaultYangActiveLevelLogModel{
		conn:  conn,
		table: "`yang_active_level_log`",
	}
}

func (m *defaultYangActiveLevelLogModel) Sharding(sharding model.ISharding) *defaultYangActiveLevelLogModel {
	m.sharding = sharding
	return m
}

func (m *defaultYangActiveLevelLogModel) Builder(tx *gorm.DB) *gorm.DB {
	return m.scopes(tx).Model(&YangActiveLevelLog{})
}

func (m *defaultYangActiveLevelLogModel) scopes(tx *gorm.DB) *gorm.DB {
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

func (m *defaultYangActiveLevelLogModel) Insert(ctx context.Context, tx *gorm.DB, data *YangActiveLevelLog) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(&data).Error
	return err
}

func (m *defaultYangActiveLevelLogModel) FindOne(ctx context.Context, id uint64) (*YangActiveLevelLog, error) {
	var resp YangActiveLevelLog
	err := m.scopes(nil).WithContext(ctx).Model(&YangActiveLevelLog{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case mysql.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultYangActiveLevelLogModel) Update(ctx context.Context, tx *gorm.DB, data *YangActiveLevelLog) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(data).Error
	return err
}

func (m *defaultYangActiveLevelLogModel) Delete(ctx context.Context, tx *gorm.DB, id uint64) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Delete(&YangActiveLevelLog{}, id).Error

	return err
}

func (m *defaultYangActiveLevelLogModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.scopes(nil).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var db = m.scopes(tx)
		return fn(db)
	})
}