package logic

import (
	"chenxi/app/mqueue/nats/internal/svc"
	"chenxi/pkg/cpb"
	"chenxi/pkg/global/cacheKey"
	"chenxi/pkg/global/constants"
	"chenxi/pkg/model"
	"chenxi/pkg/model/cloudmodel"
	"chenxi/pkg/utils"
	"context"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/proto"
)

type activeLevel struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActiveLevel(ctx context.Context, svcCtx *svc.ServiceContext) *activeLevel {
	return &activeLevel{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *activeLevel) HandleMessage(m *nats.Msg) error {
	in := &cpb.ActiveLevelEvent{}
	err := proto.Unmarshal(m.Data, in)
	if err != nil {
		return err
	}
	switch in.Event {
	case cpb.ActiveLevelEnum_reward_source:
		l.addWinPrizeRewardSource(in)
	case cpb.ActiveLevelEnum_login:
		return l.loginHandler(in)
	}
	return nil
}

func (l *activeLevel) addWinPrizeRewardSource(in *cpb.ActiveLevelEvent) {

}

// loginHandler handles the login event for the activeLevel struct.
//
// It takes an ActiveLevelEvent pointer as input parameter.
// It returns an error.
func (l *activeLevel) loginHandler(in *cpb.ActiveLevelEvent) error {
	if !l.loginCheck(in) {
		return nil
	}
	var userLevel cloudmodel.YangUserLevel
	err := l.svcCtx.CloudModel.UsersLevelModel.Builder(nil).
		Select("pid").
		Where("user_id = ? and level = ?", in.SourceUid, 1).
		Find(&userLevel).Error
	if err != nil {
		l.Logger.WithContext(l.ctx).WithFields(logx.Field("user_id", in.SourceUid)).Error(err)
		return err
	}
	var parent cloudmodel.YangUsers
	l.svcCtx.CloudModel.UsersModel.Builder(nil).Select("id").
		Where("id = ? and is_certification = ?", userLevel.Pid, 1).
		Find(&parent)
	if parent.Id == 0 {
		return nil
	}
	resultRate := l.getResultRate(float64(in.Rate), uint32(parent.Id), in.Ymd)
	var activeLevelLog cloudmodel.YangActiveLevelLog
	activeLevelLog.Uid = parent.Id
	activeLevelLog.SourceUid = uint64(in.SourceUid)
	activeLevelLog.Event = string(in.Event.Descriptor().Name())
	activeLevelLog.Rate = float64(in.Rate)
	activeLevelLog.ResultRate = resultRate
	activeLevelLog.LinkId = in.LinkId
	activeLevelLog.Pm = int64(in.Pm)
	activeLevelLog.Remark = in.Remark
	activeLevelLog.Ymd = in.Ymd
	activeLevelLog.CreateTime = uint64(time.Now().Unix())
	activeLevelLog.UpdateTime = uint64(time.Now().Unix())
	err = l.svcCtx.CloudModel.ActiveLevelLogModel.Sharding(model.NewMonthAndWeekSharing(time.Now())).
		Insert(l.ctx, nil, &activeLevelLog)
	if err != nil {
		l.Logger.WithContext(l.ctx).Error(errorx.Wrap(err, "新增活跃等级日志失败"))
		return err
	}
	return nil
}

// LoginCheck checks if a user is logged in for the active level.
//
// It takes an ActiveLevelEvent pointer as a parameter.
// It returns a boolean indicating whether the user is logged in or not.
func (l *activeLevel) loginCheck(op *cpb.ActiveLevelEvent) bool {
	var key = cacheKey.LoginActiveLevelCheckKey(op)
	cache, err := l.svcCtx.Redis.Get(key)
	if err != nil {
		l.Logger.WithContext(l.ctx).Error("redis get error")
		return false
	}
	if cache != "" {
		return false
	}
	var now = time.Now()
	var tomorrow = utils.TomorrowZeroTime()
	var expire = tomorrow.Unix() - now.Unix()
	l.svcCtx.Redis.SetnxEx(key, "1", int(expire))
	return true
}

// getResultRate calculates the result rate for a given rate, member ID, and date.
//
// Parameters:
//   - rate: the rate to calculate the result rate.
//   - memberId: the ID of the member.
//   - ymd: the date in the format "Ymd".
//
// Return:
//   - float64: the result rate.
func (l *activeLevel) getResultRate(rate float64, memberId uint32, ymd string) float64 {
	var key = cacheKey.UserRateKey(ymd)
	var uid = strconv.Itoa(int(memberId))
	result, err := l.svcCtx.Redis.Hget(key, uid)
	if err != nil {
		l.Logger.WithContext(l.ctx).Error("redis get error")
		return constants.UserWinPrizeRate_Min
	}
	if result == "" {
		return constants.UserWinPrizeRate_Min
	}

	rateConv, err := strconv.Atoi(result)
	if err != nil {
		return constants.UserWinPrizeRate_Min
	}
	currentRate := utils.DecimalDefault(rate)
	lastRate := utils.DecimalDefault(float64(rateConv))

	finalRate := lastRate.Add(currentRate).Round(2).InexactFloat64()
	if finalRate > constants.UserWinPrizeRate_Max {
		finalRate = constants.UserWinPrizeRate_Max
	}
	l.svcCtx.Redis.Hset(key, uid, strconv.FormatFloat(finalRate, 'f', -1, 64))

	var now = time.Now()
	var expireDate = now.AddDate(0, 0, 3)
	var expire = time.Date(expireDate.Year(), expireDate.Month(), expireDate.Day(), 0, 0, 0, 0, now.Location()).Unix() - now.Unix()
	l.svcCtx.Redis.Expire(key, int(expire))
	return finalRate
}
