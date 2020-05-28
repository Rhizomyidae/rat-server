package hello

import (
	"github.com/Rhizomyidae/rat-server/lib/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
}

func (c *Controller) Index(r *ghttp.Request) {
	response.JsonExit(r, 1, "success", "hello world")
}
