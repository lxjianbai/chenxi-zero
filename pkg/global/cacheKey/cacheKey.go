package cacheKey

import (
	"chenxi/pkg/cpb"
	"fmt"
)

var AppCachePrefix = "app:"

var appCacheKey = func(key string) string {
	return AppCachePrefix + key
}

func LoginActiveLevelCheckKey(op *cpb.ActiveLevelEvent) string {
	var key = fmt.Sprintf("%sbeansActive:memberId%d:event:%s:ymd:%s", AppCachePrefix, op.SourceUid, op.Event.Descriptor().Name(), op.Ymd)
	return appCacheKey(key)
}

func UserRateKey(ymd string) string {
	var key = fmt.Sprintf("%sresult_rate:%s", AppCachePrefix, ymd)
	return appCacheKey(key)
}
