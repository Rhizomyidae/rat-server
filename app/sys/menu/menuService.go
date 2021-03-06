package menu

import "github.com/Rhizomyidae/rat-server/app"

func (m *Menu) FindAllMenu(user_id int) (res []Menu) {
	// 获取取指page，指定pagesize的记录
	//err = query.Select("role_id").Where(&UserRole{UserId: user_id}).Find(&res).Error

	//_ = app.Db().Table("rat_menu m").
	//	LeftJoin("rat_role_menu rm", "m.id = rm.menu_id").
	//	LeftJoin("rat_user_role ur ", "ur.role_id = rm.role_id").
	//	Fields("m.*").
	//	Where("ur.user_id = ?", user_id).Order("m.order_num asc").Group("m.id").Structs(&res)

	sql := "select m.* from rat_menu m " +
		"left join rat_role_menu rm on m.id = rm.menu_id " +
		"left join rat_user_role ur on ur.role_id = rm.role_id " +
		"where ur.user_id = 1 " +
		"group by m.id " +
		"order by m.order_num asc "

	all, _ := app.Db().GetAll(sql)
	if len(all) > 0 {
		_ = all.Structs(&res)
		return
	}

	return
}

func constructMenuTrees(menus []Menu, parentId int, filters bool) []MenuItem {

	branch := make([]MenuItem, 0)

	for _, menu := range menus {
		if menu.ParentId == parentId {
			childList := constructMenuTrees(menus, menu.Id, filters)

			child := MenuItem{
				Id:             menu.Id,
				MenuName:       menu.MenuName,
				OrderNum:       menu.OrderNum,
				MenuType:       menu.MenuType,
				Visible:        menu.Visible,
				CreateBy:       menu.CreateBy,
				CreatedAt:      menu.CreatedAt,
				UpdateBy:       menu.UpdateBy,
				Icon:           menu.Icon,
				Component:      menu.Component,
				UpdatedAt:      menu.UpdatedAt,
				IsFrame:        menu.IsFrame,
				Perms:          menu.Perms,
				Remark:         menu.Remark,
				Url:            menu.Url,
				ParentId:       menu.ParentId,
				RoutePath:      menu.RoutePath,
				RouteName:      menu.RouteName,
				RouteComponent: menu.RouteComponent,
				RouteCache:     menu.RouteCache,
				ChildrenList:   childList,
			}
			branch = append(branch, child)
		}
	}

	return branch
}
