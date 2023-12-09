package cloudmodel

import (
	"gorm.io/gorm"
)

var _ YangCurrencyUserModel = (*customYangCurrencyUserModel)(nil)

type (
	// YangCurrencyUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangCurrencyUserModel.
	YangCurrencyUserModel interface {
		yangCurrencyUserModel
		customYangCurrencyUserLogicModel
	}

	customYangCurrencyUserModel struct {
		*defaultYangCurrencyUserModel
	}

	customYangCurrencyUserLogicModel interface {
	}
)

// NewYangCurrencyUserModel returns a model for the database table.
func NewYangCurrencyUserModel(conn *gorm.DB) YangCurrencyUserModel {
	return &customYangCurrencyUserModel{
		defaultYangCurrencyUserModel: newYangCurrencyUserModel(conn),
	}
}

func (m *defaultYangCurrencyUserModel) customCacheKeys(data *YangCurrencyUser) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
