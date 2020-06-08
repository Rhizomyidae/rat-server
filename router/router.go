package router

import (
	"github.com/Rhizomyidae/rat-server/app/cms/home"
	"github.com/Rhizomyidae/rat-server/app/hello"
	"github.com/Rhizomyidae/rat-server/app/sys/captcha"
	"github.com/Rhizomyidae/rat-server/app/sys/configs"
	"github.com/Rhizomyidae/rat-server/app/sys/dept"
	"github.com/Rhizomyidae/rat-server/app/sys/dict"
	"github.com/Rhizomyidae/rat-server/app/sys/dictData"
	"github.com/Rhizomyidae/rat-server/app/sys/menu"
	"github.com/Rhizomyidae/rat-server/app/sys/post"
	"github.com/Rhizomyidae/rat-server/app/sys/role"
	"github.com/Rhizomyidae/rat-server/app/sys/upload"
	"github.com/Rhizomyidae/rat-server/app/sys/user"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	//magicToken := r.Header.Get("Authorization")
	//if magicToken == "" {
	//	response.JsonExit(r, common.TOKEN_ERR, common.TOKEN_ERR_MSG)
	//}
	//token, b := util.ParseMagicToken(magicToken)
	//
	//if token == "" {
	//	response.JsonExit(r, common.TOKEN_ERR, common.TOKEN_VALIDE_ERR_MSG)
	//}
	//if !b {
	//	response.JsonExit(r, common.TOKEN_ERR, common.TOKEN_VALIDE_ERR_MSG)
	//}
	r.Middleware.Next()
}

// 用于配置初始化.
func init() {
	s := g.Server()

	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(MiddlewareCORS)
		s.Group("/sys", func(group *ghttp.RouterGroup) {
			group.ALL("/hello", new(hello.Controller))
			group.ALL("/user", new(user.Controller))
			group.ALL("/captcha", new(captcha.Controller))
			group.ALL("/configs", new(configs.Controller))
			group.ALL("/post", new(post.Controller))
			group.ALL("/menu", new(menu.Controller))
			group.ALL("/dict", new(dict.Controller))
			group.ALL("/dictData", new(dictData.Controller))
			group.ALL("/dept", new(dept.Controller))
			group.ALL("/role", new(role.Controller))
			group.ALL("/upload", new(upload.Controller))
		})

		s.Group("/cms", func(group *ghttp.RouterGroup) {
			group.ALL("/home", new(home.Controller))
		})
	})
}
