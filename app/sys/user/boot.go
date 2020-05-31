package user

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

var db gdb.DB

func init() {
	db = g.DB()
}

//定义控制器
type Controller struct {
}

// 注册输入参数
type SignUpInput struct {
	Username  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	Nickname  string
}
type SignUpRequest struct {
	SignUpInput
}

//登录
type SignInRequest struct {
	Username string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}
