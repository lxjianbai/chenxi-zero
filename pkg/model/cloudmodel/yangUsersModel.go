package cloudmodel

import (
	"chenxi/pkg/dao/mysql"
	"context"

	"gorm.io/gorm"
)

var _ YangUsersModel = (*customYangUsersModel)(nil)

type (
	// YangUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customYangUsersModel.
	YangUsersModel interface {
		yangUsersModel
		customYangUsersLogicModel
	}

	customYangUsersModel struct {
		*defaultYangUsersModel
	}

	customYangUsersLogicModel interface {
		FindOneByPhone(ctx context.Context, field string, phone string) (*YangUsers, error)
	}
)

// NewYangUsersModel returns a model for the database table.
func NewYangUsersModel(conn *gorm.DB) YangUsersModel {
	return &customYangUsersModel{
		defaultYangUsersModel: newYangUsersModel(conn),
	}
}

func (m *defaultYangUsersModel) customCacheKeys(data *YangUsers) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultYangUsersModel) FindOneByPhone(ctx context.Context, field string, phone string) (*YangUsers, error) {
	var resp YangUsers
	err := m.scopes(nil).WithContext(ctx).
		Model(&YangUsers{}).
		Select(field).
		Where("`phone` = ?", phone).Take(&resp).Error
	switch err {
	case nil:
		return &resp, nil
	case mysql.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
