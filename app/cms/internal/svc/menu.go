package svc

import (
	"chenxi/app/cms/internal/types"
	"chenxi/pkg/model/gamemodel"
)

func TreeMenu(menus []*gamemodel.SysMenu, parentId uint64) []*types.Menu {
	var result []*types.Menu
	for _, m := range menus {
		if m.Pid == parentId {
			var node = &types.Menu{
				Id:        int64(m.Id),
				Pid:       int64(m.Pid),
				Type:      uint8(m.Type),
				Path:      m.Path,
				Name:      m.Name,
				Redirect:  m.Redirect,
				Perms:     m.Perms,
				Component: m.Component,
				Sort:      m.Sort,
				Meta: types.MenuMeta{
					Icon:        m.Icon,
					Title:       m.Title,
					IsLink:      m.Link,
					IsHide:      uint8(m.IsHide),
					IsFull:      uint8(m.IsFull),
					IsAffix:     uint8(m.IsAffix),
					IsKeepAlive: uint8(m.IsKeepalive),
				},
				Children: TreeMenu(menus, m.Id),
			}
			result = append(result, node)
		}
	}
	return result
}
