package cloudmodel

import (
	"gorm.io/gorm"
)

var _ YangAccountbookTypeModel = (*customYangAccountbookTypeModel)(nil)

type (
	// YangAccountbookTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangAccountbookTypeModel.
	YangAccountbookTypeModel interface {
		yangAccountbookTypeModel
		customYangAccountbookTypeLogicModel
	}

	customYangAccountbookTypeModel struct {
		*defaultYangAccountbookTypeModel
	}

	customYangAccountbookTypeLogicModel interface {
	}
)

// NewYangAccountbookTypeModel returns a model for the database table.
func NewYangAccountbookTypeModel(conn *gorm.DB) YangAccountbookTypeModel {
	return &customYangAccountbookTypeModel{
		defaultYangAccountbookTypeModel: newYangAccountbookTypeModel(conn),
	}
}

func (m *defaultYangAccountbookTypeModel) customCacheKeys(data *YangAccountbookType) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
