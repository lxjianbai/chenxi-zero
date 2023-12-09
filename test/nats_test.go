package test

import (
	"chenxi/pkg/cpb"
	"chenxi/pkg/global/constants"
	"testing"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/protobuf"
	"github.com/zeromicro/go-zero/core/logx"
)

func TestNast(t *testing.T) {
	conn, _ := nats.Connect(nats.DefaultURL)
	nc, _ := nats.NewEncodedConn(conn, protobuf.PROTOBUF_ENCODER)
	for i := 1; i <= 1000; i++ {
		var data = &cpb.ActiveLevelEvent{
			Event:     cpb.ActiveLevelEnum_reward_source,
			SourceUid: uint32(i),
			Num:       uint32(i),
			Type:      cpb.WinPrizeRewardSourceTypeEnum_WinPrizeRewardSourceTypeEnumBean,
			Ymd:       "1111",
		}
		err := nc.Publish(constants.Queue.ActiveLevel.Subject, data)
		if err != nil {
			logx.Error("nats publish error: ", err)
		}
	}
}
