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
	yangSeizeBabyMemberStageOrderModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyMemberStageOrder) error

		FindOne(ctx context.Context, id uint64) (*YangSeizeBabyMemberStageOrder, error)

		Update(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyMemberStageOrder) error

		Delete(ctx context.Context, tx *gorm.DB, id uint64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
		Sharding(Sharding model.ISharding) *defaultYangSeizeBabyMemberStageOrderModel
		Builder(tx *gorm.DB) *gorm.DB
	}

	defaultYangSeizeBabyMemberStageOrderModel struct {
		conn     *gorm.DB
		table    string
		sharding model.ISharding
	}

	YangSeizeBabyMemberStageOrder struct {
		Id                 uint64        `gorm:"column:id"`                  // 自增id
		MemberId           sql.NullInt64 `gorm:"column:member_id"`           // 用户id
		StageNumId         sql.NullInt64 `gorm:"column:stage_num_id"`        // 期数id
		WinningNum         sql.NullInt64 `gorm:"column:winning_num"`         // 开奖起始号码
		EndWinningNum      sql.NullInt64 `gorm:"column:end_winning_num"`     // 开奖结束号码
		ParticipateNum     sql.NullInt64 `gorm:"column:participate_num"`     // 参与的云豆数量
		Status             sql.NullInt64 `gorm:"column:status"`              // 当前订单的状态:1=正常,2=结束,3=退款,4=超时退款
		CreatedAt          sql.NullTime  `gorm:"column:created_at"`          // 生成时间
		UpdatedAt          sql.NullTime  `gorm:"column:updated_at"`          // 更新时间
		WinningProbability sql.NullInt64 `gorm:"column:winning_probability"` // 中奖概率%
		IsWinning          sql.NullInt64 `gorm:"column:is_winning"`          // 是否中奖:1=中奖,3=未中奖,2=开奖进行中
		OpenWinningAt      sql.NullTime  `gorm:"column:open_winning_at"`     // 开奖时间
	}
)

func (YangSeizeBabyMemberStageOrder) TableName() string {
	return "`yang_seize_baby_member_stage_order`"
}

func newYangSeizeBabyMemberStageOrderModel(conn *gorm.DB) *defaultYangSeizeBabyMemberStageOrderModel {
	return &defaultYangSeizeBabyMemberStageOrderModel{
		conn:  conn,
		table: "`yang_seize_baby_member_stage_order`",
	}
}

func (m *defaultYangSeizeBabyMemberStageOrderModel) Sharding(sharding model.ISharding) *defaultYangSeizeBabyMemberStageOrderModel {
	m.sharding = sharding
	return m
}

func (m *defaultYangSeizeBabyMemberStageOrderModel) Builder(tx *gorm.DB) *gorm.DB {
	return m.scopes(tx).Model(&YangSeizeBabyMemberStageOrder{})
}

func (m *defaultYangSeizeBabyMemberStageOrderModel) scopes(tx *gorm.DB) *gorm.DB {
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

func (m *defaultYangSeizeBabyMemberStageOrderModel) Insert(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyMemberStageOrder) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(&data).Error
	return err
}

func (m *defaultYangSeizeBabyMemberStageOrderModel) FindOne(ctx context.Context, id uint64) (*YangSeizeBabyMemberStageOrder, error) {
	var resp YangSeizeBabyMemberStageOrder
	err := m.scopes(nil).WithContext(ctx).Model(&YangSeizeBabyMemberStageOrder{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case mysql.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultYangSeizeBabyMemberStageOrderModel) Update(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyMemberStageOrder) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(data).Error
	return err
}

func (m *defaultYangSeizeBabyMemberStageOrderModel) Delete(ctx context.Context, tx *gorm.DB, id uint64) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Delete(&YangSeizeBabyMemberStageOrder{}, id).Error

	return err
}

func (m *defaultYangSeizeBabyMemberStageOrderModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.scopes(nil).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var db = m.scopes(tx)
		return fn(db)
	})
}