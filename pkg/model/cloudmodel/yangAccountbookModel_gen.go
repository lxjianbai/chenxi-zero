// Code generated by goctl. DO NOT EDIT.

package cloudmodel

import (
	"context"
	"database/sql"
	"strings"
	"chenxi/pkg/dao/mysql"
	"chenxi/pkg/model"

	"gorm.io/gorm"
)

type (
	yangAccountbookModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *YangAccountbook) error

		FindOne(ctx context.Context, id int64) (*YangAccountbook, error)
		Update(ctx context.Context, tx *gorm.DB, data *YangAccountbook) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
		Sharding(Sharding model.ISharding) *defaultYangAccountbookModel
		Builder(tx *gorm.DB) *gorm.DB
	}

	defaultYangAccountbookModel struct {
		conn     *gorm.DB
		table    string
		sharding model.ISharding
	}

	YangAccountbook struct {
		Id         int64           `gorm:"column:id"`           // ID
		MemberId   int64           `gorm:"column:member_id"`    // 用户id
		CurrencyId int64           `gorm:"column:currency_id"`  // 币种
		Type       int64           `gorm:"column:type"`         // 类型id
		Content    string          `gorm:"column:content"`      // 内容(存入语言包)
		NumberType int64           `gorm:"column:number_type"`  // 收入=1/支出=2
		Current    float64         `gorm:"column:current"`      // 变动前数量
		Number     float64         `gorm:"column:number"`       // 变动的数量
		After      sql.NullFloat64 `gorm:"column:after"`        // 变动后的数量
		Fee        float64         `gorm:"column:fee"`          // 手续费
		ToMemberId int64           `gorm:"column:to_member_id"` // 目标对象ID
		AddTime    int64           `gorm:"column:add_time"`     // 添加时间
		ThirdId    int64           `gorm:"column:third_id"`     // 第三方ID
		AdRemark   string          `gorm:"column:ad_remark"`    // 管理员备注
	}
)

func (YangAccountbook) TableName() string {
	return "`yang_accountbook`"
}

func newYangAccountbookModel(conn *gorm.DB) *defaultYangAccountbookModel {
	return &defaultYangAccountbookModel{
		conn:  conn,
		table: "`yang_accountbook`",
	}
}

func (m *defaultYangAccountbookModel) Sharding(sharding model.ISharding) *defaultYangAccountbookModel {
	m.sharding = sharding
	return m
}

func (m *defaultYangAccountbookModel) Builder(tx *gorm.DB) *gorm.DB {
	return m.scopes(tx).Model(&YangAccountbook{})
}

func (m *defaultYangAccountbookModel) scopes(tx *gorm.DB) *gorm.DB {
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

func (m *defaultYangAccountbookModel) Insert(ctx context.Context, tx *gorm.DB, data *YangAccountbook) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(&data).Error
	return err
}

func (m *defaultYangAccountbookModel) FindOne(ctx context.Context, id int64) (*YangAccountbook, error) {
	var resp YangAccountbook
	err := m.scopes(nil).WithContext(ctx).Model(&YangAccountbook{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case mysql.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultYangAccountbookModel) Update(ctx context.Context, tx *gorm.DB, data *YangAccountbook) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(data).Error
	return err
}

func (m *defaultYangAccountbookModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Delete(&YangAccountbook{}, id).Error

	return err
}

func (m *defaultYangAccountbookModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.scopes(nil).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var db = m.scopes(tx)
		return fn(db)
	})
}