package cloudmodel

import (
	"gorm.io/gorm"
)

var _ YangSeizeBabyStageNumModel = (*customYangSeizeBabyStageNumModel)(nil)

type (
	// YangSeizeBabyStageNumModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangSeizeBabyStageNumModel.
	YangSeizeBabyStageNumModel interface {
		yangSeizeBabyStageNumModel
		customYangSeizeBabyStageNumLogicModel
	}

	customYangSeizeBabyStageNumModel struct {
		*defaultYangSeizeBabyStageNumModel
	}

	customYangSeizeBabyStageNumLogicModel interface {
	}
)

// NewYangSeizeBabyStageNumModel returns a model for the database table.
func NewYangSeizeBabyStageNumModel(conn *gorm.DB) YangSeizeBabyStageNumModel {
	return &customYangSeizeBabyStageNumModel{
		defaultYangSeizeBabyStageNumModel: newYangSeizeBabyStageNumModel(conn),
	}
}

func (m *defaultYangSeizeBabyStageNumModel) customCacheKeys(data *YangSeizeBabyStageNum) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
