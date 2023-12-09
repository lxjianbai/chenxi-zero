package model

import (
	"chenxi/pkg/utils"
	"fmt"
	"math"
	"strconv"
	"time"
)

type ISharding interface {
	GetTableSuffix() string
}

type ModAndMonth struct {
	ShardingVal uint32
	ShardingNum uint32
}

func NewModAndMonthSharding(shardingVal uint32, shardingNum uint32) *ModAndMonth {
	return &ModAndMonth{
		ShardingVal: shardingVal,
		ShardingNum: shardingNum,
	}
}

func (mm *ModAndMonth) GetTableSuffix() string {
	var seq = mm.ShardingVal % mm.ShardingNum
	return fmt.Sprintf("%s_%d", time.Now().Format("200601"), seq)
}

type MonthAndWeek struct {
	time time.Time
}

func NewMonthAndWeekSharing(time time.Time) *MonthAndWeek {
	return &MonthAndWeek{
		time: time,
	}
}

func (mw *MonthAndWeek) GetTableSuffix() string {
	var ym = mw.time.Format(utils.MonthOnly_Ym)
	d, _ := strconv.Atoi(mw.time.Format(utils.DateOnly_d))
	return fmt.Sprintf("%sw%d", ym, int(math.Ceil(float64(d)/7)))
}
