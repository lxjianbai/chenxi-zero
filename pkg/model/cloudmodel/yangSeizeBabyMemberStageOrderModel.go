package cloudmodel

import (
	"gorm.io/gorm"
)

var _ YangSeizeBabyMemberStageOrderModel = (*customYangSeizeBabyMemberStageOrderModel)(nil)

type (
	// YangSeizeBabyMemberStageOrderModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangSeizeBabyMemberStageOrderModel.
	YangSeizeBabyMemberStageOrderModel interface {
		yangSeizeBabyMemberStageOrderModel
		customYangSeizeBabyMemberStageOrderLogicModel
	}

	customYangSeizeBabyMemberStageOrderModel struct {
		*defaultYangSeizeBabyMemberStageOrderModel
	}

	customYangSeizeBabyMemberStageOrderLogicModel interface {
	}
)

// NewYangSeizeBabyMemberStageOrderModel returns a model for the database table.
func NewYangSeizeBabyMemberStageOrderModel(conn *gorm.DB) YangSeizeBabyMemberStageOrderModel {
	return &customYangSeizeBabyMemberStageOrderModel{
		defaultYangSeizeBabyMemberStageOrderModel: newYangSeizeBabyMemberStageOrderModel(conn),
	}
}

func (m *defaultYangSeizeBabyMemberStageOrderModel) customCacheKeys(data *YangSeizeBabyMemberStageOrder) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
