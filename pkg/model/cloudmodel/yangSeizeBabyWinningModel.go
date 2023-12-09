package cloudmodel

import (
	"gorm.io/gorm"
)

var _ YangSeizeBabyWinningModel = (*customYangSeizeBabyWinningModel)(nil)

type (
	// YangSeizeBabyWinningModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangSeizeBabyWinningModel.
	YangSeizeBabyWinningModel interface {
		yangSeizeBabyWinningModel
		customYangSeizeBabyWinningLogicModel
	}

	customYangSeizeBabyWinningModel struct {
		*defaultYangSeizeBabyWinningModel
	}

	customYangSeizeBabyWinningLogicModel interface {
	}
)

// NewYangSeizeBabyWinningModel returns a model for the database table.
func NewYangSeizeBabyWinningModel(conn *gorm.DB) YangSeizeBabyWinningModel {
	return &customYangSeizeBabyWinningModel{
		defaultYangSeizeBabyWinningModel: newYangSeizeBabyWinningModel(conn),
	}
}

func (m *defaultYangSeizeBabyWinningModel) customCacheKeys(data *YangSeizeBabyWinning) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
