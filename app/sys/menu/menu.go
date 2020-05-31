package menu

import (
	"github.com/Rhizomyidae/rat-server/lib/response"
	"github.com/gogf/gf/net/ghttp"
)

func (c *Controller) FindAllMenu(r *ghttp.Request) {
	model := NewMenu()
	menuData := model.FindAllMenu(0)
	dataMap := constructMenuTrees(menuData, 0, true)
	response.JsonOK(r, dataMap)
}
