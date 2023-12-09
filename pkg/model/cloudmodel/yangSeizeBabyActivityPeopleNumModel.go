package cloudmodel

import (
	"gorm.io/gorm"
)

var _ YangSeizeBabyActivityPeopleNumModel = (*customYangSeizeBabyActivityPeopleNumModel)(nil)

type (
	// YangSeizeBabyActivityPeopleNumModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangSeizeBabyActivityPeopleNumModel.
	YangSeizeBabyActivityPeopleNumModel interface {
		yangSeizeBabyActivityPeopleNumModel
		customYangSeizeBabyActivityPeopleNumLogicModel
	}

	customYangSeizeBabyActivityPeopleNumModel struct {
		*defaultYangSeizeBabyActivityPeopleNumModel
	}

	customYangSeizeBabyActivityPeopleNumLogicModel interface {
	}
)

// NewYangSeizeBabyActivityPeopleNumModel returns a model for the database table.
func NewYangSeizeBabyActivityPeopleNumModel(conn *gorm.DB) YangSeizeBabyActivityPeopleNumModel {
	return &customYangSeizeBabyActivityPeopleNumModel{
		defaultYangSeizeBabyActivityPeopleNumModel: newYangSeizeBabyActivityPeopleNumModel(conn),
	}
}

func (m *defaultYangSeizeBabyActivityPeopleNumModel) customCacheKeys(data *YangSeizeBabyActivityPeopleNum) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
