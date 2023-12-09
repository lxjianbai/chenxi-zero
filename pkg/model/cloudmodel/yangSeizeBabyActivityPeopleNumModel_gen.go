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
	yangSeizeBabyActivityPeopleNumModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyActivityPeopleNum) error

		FindOne(ctx context.Context, id uint64) (*YangSeizeBabyActivityPeopleNum, error)
		Update(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyActivityPeopleNum) error

		Delete(ctx context.Context, tx *gorm.DB, id uint64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
		Sharding(Sharding model.ISharding) *defaultYangSeizeBabyActivityPeopleNumModel
		Builder(tx *gorm.DB) *gorm.DB
	}

	defaultYangSeizeBabyActivityPeopleNumModel struct {
		conn     *gorm.DB
		table    string
		sharding model.ISharding
	}

	YangSeizeBabyActivityPeopleNum struct {
		Id           uint64        `gorm:"column:id"`            // 自增id
		MemberId     sql.NullInt64 `gorm:"column:member_id"`     // 用户id
		ActivityId   sql.NullInt64 `gorm:"column:activity_id"`   // 活动id
		FrequencyNum sql.NullInt64 `gorm:"column:frequency_num"` // 参与活动的人次数
		CreatedAt    sql.NullTime  `gorm:"column:created_at"`    // 开始参与活动的时间
		UpdatedAt    sql.NullTime  `gorm:"column:updated_at"`    // 最后参与活动的时间
		IsWinning    sql.NullInt64 `gorm:"column:is_winning"`    // 是否中奖:1=中奖,3=未中奖,2=开奖进行中
	}
)

func (YangSeizeBabyActivityPeopleNum) TableName() string {
	return "`yang_seize_baby_activity_people_num`"
}

func newYangSeizeBabyActivityPeopleNumModel(conn *gorm.DB) *defaultYangSeizeBabyActivityPeopleNumModel {
	return &defaultYangSeizeBabyActivityPeopleNumModel{
		conn:  conn,
		table: "`yang_seize_baby_activity_people_num`",
	}
}

func (m *defaultYangSeizeBabyActivityPeopleNumModel) Sharding(sharding model.ISharding) *defaultYangSeizeBabyActivityPeopleNumModel {
	m.sharding = sharding
	return m
}

func (m *defaultYangSeizeBabyActivityPeopleNumModel) Builder(tx *gorm.DB) *gorm.DB {
	return m.scopes(tx).Model(&YangSeizeBabyActivityPeopleNum{})
}

func (m *defaultYangSeizeBabyActivityPeopleNumModel) scopes(tx *gorm.DB) *gorm.DB {
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

func (m *defaultYangSeizeBabyActivityPeopleNumModel) Insert(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyActivityPeopleNum) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(&data).Error
	return err
}

func (m *defaultYangSeizeBabyActivityPeopleNumModel) FindOne(ctx context.Context, id uint64) (*YangSeizeBabyActivityPeopleNum, error) {
	var resp YangSeizeBabyActivityPeopleNum
	err := m.scopes(nil).WithContext(ctx).Model(&YangSeizeBabyActivityPeopleNum{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case mysql.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultYangSeizeBabyActivityPeopleNumModel) Update(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyActivityPeopleNum) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(data).Error
	return err
}

func (m *defaultYangSeizeBabyActivityPeopleNumModel) Delete(ctx context.Context, tx *gorm.DB, id uint64) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Delete(&YangSeizeBabyActivityPeopleNum{}, id).Error

	return err
}

func (m *defaultYangSeizeBabyActivityPeopleNumModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.scopes(nil).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var db = m.scopes(tx)
		return fn(db)
	})
}