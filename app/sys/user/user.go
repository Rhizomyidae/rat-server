package user

import (
	"github.com/Rhizomyidae/rat-server/common"
	"github.com/Rhizomyidae/rat-server/lib/response"
	"github.com/Rhizomyidae/rat-server/util"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

func (c *Controller) Login(r *ghttp.Request) {
	var data *SignInRequest
	if err := r.Parse(&data); err != nil {
		response.JsonERR(r, err.Error())
	}
	if one, err := signIn(data.Username, data.Password); err != nil {
		response.JsonERR(r, err.Error())
	} else {
		encryptString, _ := gmd5.EncryptString(gconv.String(one.GMap().Get("user_name")))
		str := util.CreateMagicToken(g.Map{
			"id":           one.GMap().Get("id"),
			"verification": encryptString,
			"user_name":    one.GMap().Get("user_name"),
		})

		if str != "" {
			urlimg := one.GMap().Get("url")
			if urlimg == "" || urlimg == nil {
				urlimg = "http://zskj-app.oss-cn-hangzhou.aliyuncs.com/%E7%94%B5%E7%81%AF%E6%B3%A1.jpg"
			}
			response.JsonOK(r, g.Map{
				"token":    str,
				"userId":   one.GMap().Get("id"),
				"userName": one.GMap().Get("user_name"),
				"nickname": one.GMap().Get("nickname"),
			})

			signInRecord(one.GMap().Get("id"), r)
		} else {
			response.JsonERR(r, "登录失败")
		}
	}
}

func (c *Controller) CheckToken(r *ghttp.Request) {
	token := r.Header.Get("Authorization")

	b, message, code := util.CheckToken(token)

	if !b {
		response.JsonERR(r, message)
	}

	jsonData := make(map[string]interface{}, 1)
	jsonData["user_id"] = code
	jsonData["user_name"] = common.UserName

	response.JsonOK(r, jsonData)
}

func (c *Controller) Logout(r *ghttp.Request) {
	response.JsonOK(r)
}
