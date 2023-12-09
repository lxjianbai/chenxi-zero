package cloudmodel

import (
	"gorm.io/gorm"
)

var _ YangSeizeBabyGoodsModel = (*customYangSeizeBabyGoodsModel)(nil)

type (
	// YangSeizeBabyGoodsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangSeizeBabyGoodsModel.
	YangSeizeBabyGoodsModel interface {
		yangSeizeBabyGoodsModel
		customYangSeizeBabyGoodsLogicModel
	}

	customYangSeizeBabyGoodsModel struct {
		*defaultYangSeizeBabyGoodsModel
	}

	customYangSeizeBabyGoodsLogicModel interface {
	}
)

// NewYangSeizeBabyGoodsModel returns a model for the database table.
func NewYangSeizeBabyGoodsModel(conn *gorm.DB) YangSeizeBabyGoodsModel {
	return &customYangSeizeBabyGoodsModel{
		defaultYangSeizeBabyGoodsModel: newYangSeizeBabyGoodsModel(conn),
	}
}

func (m *defaultYangSeizeBabyGoodsModel) customCacheKeys(data *YangSeizeBabyGoods) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
