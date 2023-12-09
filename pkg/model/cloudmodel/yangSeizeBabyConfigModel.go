package cloudmodel

import (
	"gorm.io/gorm"
)

var _ YangSeizeBabyConfigModel = (*customYangSeizeBabyConfigModel)(nil)

type (
	// YangSeizeBabyConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangSeizeBabyConfigModel.
	YangSeizeBabyConfigModel interface {
		yangSeizeBabyConfigModel
		customYangSeizeBabyConfigLogicModel
	}

	customYangSeizeBabyConfigModel struct {
		*defaultYangSeizeBabyConfigModel
	}

	customYangSeizeBabyConfigLogicModel interface {
	}
)

// NewYangSeizeBabyConfigModel returns a model for the database table.
func NewYangSeizeBabyConfigModel(conn *gorm.DB) YangSeizeBabyConfigModel {
	return &customYangSeizeBabyConfigModel{
		defaultYangSeizeBabyConfigModel: newYangSeizeBabyConfigModel(conn),
	}
}

func (m *defaultYangSeizeBabyConfigModel) customCacheKeys(data *YangSeizeBabyConfig) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
