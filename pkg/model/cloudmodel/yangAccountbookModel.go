package cloudmodel

import (
	"gorm.io/gorm"
)

var _ YangAccountbookModel = (*customYangAccountbookModel)(nil)

type (
	// YangAccountbookModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangAccountbookModel.
	YangAccountbookModel interface {
		yangAccountbookModel
		customYangAccountbookLogicModel
	}

	customYangAccountbookModel struct {
		*defaultYangAccountbookModel
	}

	customYangAccountbookLogicModel interface {
	}
)

// NewYangAccountbookModel returns a model for the database table.
func NewYangAccountbookModel(conn *gorm.DB) YangAccountbookModel {
	return &customYangAccountbookModel{
		defaultYangAccountbookModel: newYangAccountbookModel(conn),
	}
}

func (m *defaultYangAccountbookModel) customCacheKeys(data *YangAccountbook) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
