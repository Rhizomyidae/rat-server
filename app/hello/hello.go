package hello

import (
	"github.com/Rhizomyidae/rat-server/app"
	"github.com/Rhizomyidae/rat-server/lib/response"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	app.Controller
}

func (c *Controller) Index(r *ghttp.Request) {
	response.SuccessResult(r, "hello word")
}
