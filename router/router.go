package router

import "github.com/gogf/gf/frame/g"

// 用于配置初始化.
func init() {
	// Avoid 404 for some browsers requesting "/favicon.ico".
	g.Server().SetRewrite("/favicon.ico", "/resource/image/favicon.ico")
}
