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
	yangSeizeBabyStageNumModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyStageNum) error

		FindOne(ctx context.Context, id uint64) (*YangSeizeBabyStageNum, error)
		Update(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyStageNum) error

		Delete(ctx context.Context, tx *gorm.DB, id uint64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
		Sharding(Sharding model.ISharding) *defaultYangSeizeBabyStageNumModel
		Builder(tx *gorm.DB) *gorm.DB
	}

	defaultYangSeizeBabyStageNumModel struct {
		conn     *gorm.DB
		table    string
		sharding model.ISharding
	}

	YangSeizeBabyStageNum struct {
		Id                   uint64          `gorm:"column:id"`                     // id
		StageName            sql.NullString  `gorm:"column:stage_name"`             // 名称
		Stageimage           sql.NullString  `gorm:"column:stageimage"`             // 图片
		CoinUnitPrice        sql.NullFloat64 `gorm:"column:coin_unit_price"`        // 云豆单价
		TotalNeedPeopleNum   sql.NullInt64   `gorm:"column:total_need_people_num"`  // 总共需要人数
		SurplusPeopleNum     sql.NullInt64   `gorm:"column:surplus_people_num"`     // 剩余人数
		ParticipatePeopleNum sql.NullInt64   `gorm:"column:participate_people_num"` // 当前参与人数
		LeastNum             sql.NullInt64   `gorm:"column:least_num"`              // 最少参与的云豆
		UpdatedAt            sql.NullTime    `gorm:"column:updated_at"`             // 更新时间
		CreatedAt            sql.NullTime    `gorm:"column:created_at"`             // 生成时间
		EndAt                sql.NullTime    `gorm:"column:end_at"`                 // 结束时间
		StageNumType         sql.NullInt64   `gorm:"column:stage_num_type"`         // 类型:1=云豆夺宝,2=实物夺宝
		GoodsId              sql.NullInt64   `gorm:"column:goods_id"`               // 关联的商品id
		WinningNum           sql.NullInt64   `gorm:"column:winning_num"`            // 中奖号码
		StageStatus          sql.NullInt64   `gorm:"column:stage_status"`           // 当前期数的状态:1=正常,2=结束,3=超时结束
		MemberId             sql.NullInt64   `gorm:"column:member_id"`              // 用户id-- 发布的用户
		WinningUserId        sql.NullInt64   `gorm:"column:winning_user_id"`        // 中奖用户的id
		AppointWinningUser   int64           `gorm:"column:appoint_winning_user"`   // 指定的中奖用户
	}
)

func (YangSeizeBabyStageNum) TableName() string {
	return "`yang_seize_baby_stage_num`"
}

func newYangSeizeBabyStageNumModel(conn *gorm.DB) *defaultYangSeizeBabyStageNumModel {
	return &defaultYangSeizeBabyStageNumModel{
		conn:  conn,
		table: "`yang_seize_baby_stage_num`",
	}
}

func (m *defaultYangSeizeBabyStageNumModel) Sharding(sharding model.ISharding) *defaultYangSeizeBabyStageNumModel {
	m.sharding = sharding
	return m
}

func (m *defaultYangSeizeBabyStageNumModel) Builder(tx *gorm.DB) *gorm.DB {
	return m.scopes(tx).Model(&YangSeizeBabyStageNum{})
}

func (m *defaultYangSeizeBabyStageNumModel) scopes(tx *gorm.DB) *gorm.DB {
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

func (m *defaultYangSeizeBabyStageNumModel) Insert(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyStageNum) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(&data).Error
	return err
}

func (m *defaultYangSeizeBabyStageNumModel) FindOne(ctx context.Context, id uint64) (*YangSeizeBabyStageNum, error) {
	var resp YangSeizeBabyStageNum
	err := m.scopes(nil).WithContext(ctx).Model(&YangSeizeBabyStageNum{}).Where("`id` = ?", id).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case mysql.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultYangSeizeBabyStageNumModel) Update(ctx context.Context, tx *gorm.DB, data *YangSeizeBabyStageNum) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Save(data).Error
	return err
}

func (m *defaultYangSeizeBabyStageNumModel) Delete(ctx context.Context, tx *gorm.DB, id uint64) error {
	db := m.scopes(tx)
	err := db.WithContext(ctx).Delete(&YangSeizeBabyStageNum{}, id).Error

	return err
}

func (m *defaultYangSeizeBabyStageNumModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.scopes(nil).WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var db = m.scopes(tx)
		return fn(db)
	})
}