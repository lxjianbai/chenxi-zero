package services

import (
	"chenxi/pkg/cpb"
	"chenxi/pkg/global/constants"
	"chenxi/pkg/model"
	"chenxi/pkg/model/cloudmodel"
	"chenxi/pkg/utils"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/gookit/goutil"
	"github.com/shopspring/decimal"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

// 用户资产操作
type accountBookServices struct {
	ctx context.Context
	logx.Logger
	cloudDb     *gorm.DB
	cloudModel  *cloudmodel.CloudModel
	lotteryType []cpb.AccountbookEnum
}

// NewCurrencyUserServices creates a new instance of the accountBookServices struct.
//
// It takes a context.Context, a logx.Logger, a *gorm.DB, and a *global.Global as parameters.
// It returns a pointer to the accountBookServices struct and an error.
func NewCurrencyUserServices(ctx context.Context, logger logx.Logger, cloudDb *gorm.DB) (*accountBookServices, error) {
	if cloudDb == nil {
		return nil, errors.New("db conn err")
	}
	service := &accountBookServices{
		ctx:        ctx,
		Logger:     logger.WithContext(ctx),
		cloudDb:    cloudDb,
		cloudModel: cloudmodel.NewCloudModel(cloudDb),
		lotteryType: []cpb.AccountbookEnum{
			cpb.AccountbookEnum_AccountbookEnumLotteryType,
		},
	}

	return service, nil
}

// 操作用户资产
func (s *accountBookServices) OperatingCurrency(tx *gorm.DB, currencyUser *cloudmodel.YangCurrencyUser, op *cpb.OperatingCurrency) error {
	if tx == nil || currencyUser == nil || op == nil {
		return errors.New("operatingCurrency parameter err")
	}

	var accountbookType cloudmodel.YangAccountbookType
	err := s.cloudModel.AccountbookTypeModel.Builder(tx).
		Select("name").
		Where("id = ?", op.TypeId).
		Find(&accountbookType).Error
	if err != nil {
		s.Logger.WithFields(logx.Field("type_id", op.TypeId)).Error(err)
		return err
	}

	var nowNum = utils.DecimalCurrency(currencyUser.Num) // 当前数量
	var opNum = utils.DecimalCurrency(op.Num)            // 操作数量
	var after decimal.Decimal
	var fee = utils.DecimalCurrency(op.Fee) // 手续费
	switch op.NumberType {
	case cpb.OperatingCurrencyEnum_OperatingCurrencyIn:
		after = nowNum.Add(opNum)
	case cpb.OperatingCurrencyEnum_OperatingCurrencyOut:
		after = nowNum.Sub(opNum)
		if op.Fee > 0 {
			after = after.Sub(fee)
		}
	}

	if decimal.Zero.Cmp(after) > 0 {
		err = errors.New("operatingCurrency insufficient balance")
		s.Logger.WithFields(logx.Field("member_id", op.MemberId), logx.Field("currency_id", op.CurrencyId)).Error(err)
		return err
	}

	if op.Fee > 0 && op.NumberType == cpb.OperatingCurrencyEnum_OperatingCurrencyOut {
		opNum = opNum.Add(fee)
	}

	var accountBook cloudmodel.YangAccountbook
	accountBook.MemberId = int64(op.MemberId)
	accountBook.CurrencyId = int64(op.CurrencyId)
	accountBook.Type = int64(op.NumberType)
	accountBook.Content = accountbookType.Name
	accountBook.Current = nowNum.InexactFloat64()
	accountBook.Number = opNum.InexactFloat64()
	accountBook.After = sql.NullFloat64{
		Float64: after.InexactFloat64(),
		Valid:   true,
	}
	accountBook.Fee = fee.InexactFloat64()
	accountBook.ThirdId = int64(op.ThirdId)
	accountBook.ToMemberId = int64(op.ToMemberId)

	// 银币不加入账本 ??? php代码
	if op.CurrencyId != uint32(cpb.CurrencyEnum_CurrencySilver) {
		if err := s.cloudModel.AccountbookModel.Sharding(model.NewModAndMonthSharding(op.MemberId, constants.ShardingNum_Ten)).
			Insert(s.ctx, tx, &accountBook); err != nil {
			s.Logger.WithFields(logx.Field("member_id", op.MemberId), logx.Field("currency_id", op.CurrencyId)).Error(err)
			return err
		}
	}

	var currencyUserOldNum = currencyUser.Num
	var currencyUserOldFrozen = currencyUser.FreezeNum
	var currencyUserFreezeNum = utils.DecimalCurrency(currencyUser.FreezeNum)

	if op.NumberType == cpb.OperatingCurrencyEnum_OperatingCurrencyIn {
		currencyUser.Num = after.InexactFloat64()
		if op.IsFrozen {
			currencyUser.FreezeNum = currencyUserFreezeNum.Sub(opNum).InexactFloat64() // 减少冻结
		}
		if goutil.Contains(s.lotteryType, op.TypeId) {
			var currencyUserLotteryNum = utils.DecimalCurrency(currencyUser.LotteryNum)
			currencyUser.LotteryNum = currencyUserLotteryNum.Add(opNum).InexactFloat64()
		}
	} else {
		currencyUser.Num = after.InexactFloat64()
		if op.IsFrozen {
			currencyUser.FreezeNum = currencyUserFreezeNum.Add(opNum.Add(fee)).InexactFloat64() // 增加冻结
		}
	}
	currencyUser.UpdatedAt = uint64(time.Now().Unix())
	if currencyUser.Num < 0 || currencyUser.FreezeNum < 0 {
		err = errors.New("operatingCurrency asset operation abnormal")
		s.Logger.WithFields(logx.Field("member_id", op.MemberId), logx.Field("currency_id", op.CurrencyId)).Error(err)
		return err
	}
	result := s.cloudModel.CurrencyUserModel.Builder(tx).
		Where("member_id = ? and num = ? and freeze_num = ?", op.MemberId, currencyUserOldNum, currencyUserOldFrozen).
		Save(&currencyUser)
	if result.Error != nil {
		s.Logger.WithFields(logx.Field("member_id", op.MemberId), logx.Field("currency_id", op.CurrencyId)).Error(err)
		return result.Error
	}
	if result.RowsAffected == 0 {
		err = errors.New("operatingCurrency update currencyUser rowsAffected zero")
		s.Logger.WithContext(s.ctx).WithFields(logx.Field("member_id", op.MemberId)).Error(err)
		return err
	}
	return nil
}
