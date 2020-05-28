package router

import (
	"github.com/Rhizomyidae/rat-server/app/hello"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 用于配置初始化.
func init() {
	s := g.Server()

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/hello", new(hello.Controller))
	})
}
