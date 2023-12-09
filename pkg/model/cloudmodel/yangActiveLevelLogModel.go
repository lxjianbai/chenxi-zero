package cloudmodel

import (
	"gorm.io/gorm"
)

var _ YangActiveLevelLogModel = (*customYangActiveLevelLogModel)(nil)

type (
	// YangActiveLevelLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangActiveLevelLogModel.
	YangActiveLevelLogModel interface {
		yangActiveLevelLogModel
		customYangActiveLevelLogLogicModel
	}

	customYangActiveLevelLogModel struct {
		*defaultYangActiveLevelLogModel
	}

	customYangActiveLevelLogLogicModel interface {
	}
)

// NewYangActiveLevelLogModel returns a model for the database table.
func NewYangActiveLevelLogModel(conn *gorm.DB) YangActiveLevelLogModel {
	return &customYangActiveLevelLogModel{
		defaultYangActiveLevelLogModel: newYangActiveLevelLogModel(conn),
	}
}

func (m *defaultYangActiveLevelLogModel) customCacheKeys(data *YangActiveLevelLog) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
