package cloudmodel

import "gorm.io/gorm"

type CloudModel struct {
	UsersModel           YangUsersModel
	UsersLevelModel      YangUserLevelModel
	ActiveLevelLogModel  YangActiveLevelLogModel
	AccountbookTypeModel YangAccountbookTypeModel
	AccountbookModel     YangAccountbookModel
	CurrencyUserModel    YangCurrencyUserModel
}

func NewCloudModel(db *gorm.DB) *CloudModel {
	return &CloudModel{
		UsersModel:           NewYangUsersModel(db),
		UsersLevelModel:      NewYangUserLevelModel(db),
		ActiveLevelLogModel:  NewYangActiveLevelLogModel(db),
		AccountbookTypeModel: NewYangAccountbookTypeModel(db),
		AccountbookModel:     NewYangAccountbookModel(db),
		CurrencyUserModel:    NewYangCurrencyUserModel(db),
	}
}
