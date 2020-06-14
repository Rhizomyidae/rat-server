package home

import (
	"github.com/Rhizomyidae/rat-server/lib/response"
	"github.com/gogf/gf/net/ghttp"
)

func (c *Controller) Index(r *ghttp.Request) {
	var data *SignInRequest
	if err := r.Parse(&data); err != nil {
		response.JsonERR(r, err.Error())
	}
	if one, err := queryArticle(); err != nil {
		response.JsonERR(r, err.Error())
	} else {
		response.JsonOK(r, one)
	}
}
