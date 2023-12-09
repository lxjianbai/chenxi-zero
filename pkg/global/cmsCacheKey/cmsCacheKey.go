package cmscachekey

import "fmt"

var CmsCachePrefix = "cms:"

var cmsCacheKey = func(key string) string {
	return CmsCachePrefix + key
}

func SysUserToken(userId uint32) string {
	var key = fmt.Sprintf("%d", userId)
	return cmsCacheKey(key)
}

func SysMenu() string {
	var key = "sys:menu"
	return cmsCacheKey(key)
}

// SysUserMenu
// SysMenu里面加了前缀，这里不能在加了
func SysUserMenu(userId uint32) string {
	return fmt.Sprintf("%s:%d", SysMenu(), userId)
}
