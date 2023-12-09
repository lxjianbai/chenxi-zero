package cloudmodel

import (
	"gorm.io/gorm"
)

var _ YangUserLevelModel = (*customYangUserLevelModel)(nil)

type (
	// YangUserLevelModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangUserLevelModel.
	YangUserLevelModel interface {
		yangUserLevelModel
		customYangUserLevelLogicModel
	}

	customYangUserLevelModel struct {
		*defaultYangUserLevelModel
	}

	customYangUserLevelLogicModel interface {
	}
)

// NewYangUserLevelModel returns a model for the database table.
func NewYangUserLevelModel(conn *gorm.DB) YangUserLevelModel {
	return &customYangUserLevelModel{
		defaultYangUserLevelModel: newYangUserLevelModel(conn),
	}
}

func (m *defaultYangUserLevelModel) customCacheKeys(data *YangUserLevel) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
