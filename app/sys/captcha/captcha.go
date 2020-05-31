package captcha

import (
	"github.com/Rhizomyidae/rat-server/app"
	"github.com/Rhizomyidae/rat-server/lib/response"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	app.Controller
}

type Captcha struct {
	Ticket  string `json:"ticket" form:"ticket"`
	Randstr string `json:"randstr"  form:"randstr"`
}

func (c *Controller) Check(r *ghttp.Request) {
	var captcha Captcha
	if err := r.Parse(&captcha); err != nil {
		response.JsonERR(r, err.Error())
	}

	Ticket := captcha.Ticket
	Randstr := captcha.Randstr
	UserIp := r.GetClientIp()

	req := gmap.New(true)
	req.Set("Action", "DescribeCaptchaResult")
	req.Set("CaptchaType", "9")
	req.Set("Ticket", Ticket)
	req.Set("UserIp", UserIp)
	req.Set("Version", "2019-07-22")
	req.Set("Randstr", Randstr)
	req.Set("CaptchaAppId", "1251180753")
	req.Set("AppSecretKey", "zlqfnkcniyxxZvJQV2I2Xona69vQFpAE")

	clientRespons, err := ghttp.Get("https://captcha.tencentcloudapi.com/", req)

	if err != nil {
		response.JsonERR(r, err.Error())
	}

	allString := clientRespons.ReadAllString()
	json, _ := gjson.DecodeToJson(allString)
	if json.GetInt("ret_code") == 0 {
		response.JsonOK(r)
	} else {
		response.JsonERR(r, allString)
	}
}

func (c *Controller) Hander(r *ghttp.Request) {
	var captcha Captcha
	if err := r.Parse(&captcha); err != nil {
		response.JsonERR(r, err.Error())
	}

	Ticket := captcha.Ticket
	Randstr := captcha.Randstr
	UserIp := r.GetClientIp()

	req := gmap.New(true)
	req.Set("Ticket", Ticket)
	req.Set("UserIP", UserIp)
	req.Set("Randstr", Randstr)
	req.Set("aid", "2076088864")
	req.Set("AppSecretKey", "06bEYSvZpRbeo6n_bMR0G_g**")

	clientRespons, err := ghttp.Get("https://ssl.captcha.qq.com/ticket/verify", req)

	if err != nil {
		response.JsonERR(r, err.Error())
	}

	allString := clientRespons.ReadAllString()
	json, _ := gjson.DecodeToJson(allString)
	if json.GetInt("response") == 1 {
		response.JsonOK(r)
	} else {
		response.JsonERR(r, allString)
	}
}
