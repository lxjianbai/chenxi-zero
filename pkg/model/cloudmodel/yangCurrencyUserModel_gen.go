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
	yangCurrencyUserModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *YangCurrencyUser) error

		FindOne(ctx context.Context, id uint64) (*YangCurrencyUser, error)
		Update(ctx context.Context, tx *gorm.DB, data *YangCurrencyUser) error

		Delete(ctx context.Context, tx *gorm.DB, id uint64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
		Sharding(Sharding model.ISharding) *defaultYangCurrencyUserModel
		Builder(tx *gorm.DB) *gorm.DB
	}

	defaultYangCurrencyUserModel struct {
		conn     *gorm.DB
		table    string
		sharding model.ISharding
	}

	YangCurrencyUser struct {
		Id         uint64        `gorm:"column:id"`
		MemberId   sql.NullInt64 `gorm:"column:member_id"`   // 用户id
		CurrencyId sql.NullInt64 `gorm:"column:currency_id"` // 资产id
		Num        float64       `gorm:"column:num"`         // 可用数量
		FreezeNum  float64       `gorm:"column:freeze_num"`  // 冻结数量
		UpdatedAt  uint64        `gorm:"column:updated_at"`  // 更新时间
		LotteryNum float64       `gorm:"column:lottery_num"` // 中奖获得的数量
	}
)

func (YangCurrencyUser) TableName() string {
	return "`yang_currency_user`"
}

func newYangCurrencyUserModel(conn *gorm.DB) *defaultYangCurrencyUserModel {
	return &defaultYangCurrencyUserModel{
		conn:  conn,
		table: "`yang_currency_user`",
	}
}

func (m *defaultYangCurrencyUserModel) Sharding(sharding model.ISharding) *defaultYangCurrencyUserModel {
	m.sharding = sharding
	return m
}

func (m *defaultYangCurrencyUserModel) Builder(tx *gorm.DB) *gorm.DB {
	return m.scopes(tx).Model(&YangCurrencyUser{})
}

func (m *defaultYangCurrencyUserModel) scopes(tx *gorm.DB) *gorm.DB {
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

func (m *defaultYangCurrencyUserModel) Insert(ctx context.Context, tx *gorm.DB, data *YangCurrencyUser) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(&data).Error
	return err
}

func (m *defaultYangCurrencyUserModel) FindOne(ctx context.Context, id uint64) (*YangCurrencyUser, error) {
	var resp YangCurrencyUser
	err := m.scopes(nil).WithContext(ctx).Model(&YangCurrencyUser{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case mysql.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultYangCurrencyUserModel) Update(ctx context.Context, tx *gorm.DB, data *YangCurrencyUser) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(data).Error
	return err
}

func (m *defaultYangCurrencyUserModel) Delete(ctx context.Context, tx *gorm.DB, id uint64) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Delete(&YangCurrencyUser{}, id).Error

	return err
}

func (m *defaultYangCurrencyUserModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.scopes(nil).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var db = m.scopes(tx)
		return fn(db)
	})
}
