package user

import (
	"errors"
	"fmt"
	"github.com/Rhizomyidae/rat-server/app"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/gvalid"
)

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func signIn(passport, password string) (gdb.Record, error) {
	rone, err := app.Db().GetOne("select * from rat_user where user_name=? limit 1", passport)
	if err != nil {
		return rone, err
	}
	if rone == nil {
		return rone, errors.New("账户名不存在")
	}

	rsalt := gconv.String(rone.GMap().Get("salt"))
	rpassword := gconv.String(rone.GMap().Get("password"))
	rs := password + rsalt
	eStr, _ := gmd5.EncryptString(rs)

	if eStr != rpassword {
		return rone, errors.New("密码不正确")
	}
	return rone, err
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func signInRecord(id interface{}, r *ghttp.Request) {
	_, _ = app.Db().Exec("update rat_user set login_date=now() and login_ip=? where id=?", r.GetClientIp(), id)
}

func signUp(data *SignUpInput) error {
	// 输入参数检查
	if e := gvalid.CheckStruct(data, nil); e != nil {
		return errors.New(e.FirstString())
	}
	// 昵称为非必需参数，默认使用账号名称
	if data.Nickname == "" {
		data.Nickname = data.Username
	}
	// 账号唯一性数据检查
	if !checkPassport(data.Username) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", data.Username))
	}
	// 昵称唯一性数据检查
	if !checkNickName(data.Nickname) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", data.Nickname))
	}

	encryptString, _ := gmd5.EncryptString(data.Password)
	_, _ = app.Db().Table("rat_user").Data(g.Map{
		"username": data.Username,
		"password": encryptString,
		"nickname": data.Nickname,
		"created":  gtime.Now(),
	}).Insert()

	return nil
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func checkPassport(passport string) bool {
	if i, err := app.Db().GetAll("select * from rat_user where user_name=?", passport); err != nil {
		return false
	} else {
		return len(i) == 0
	}
}

// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func checkNickName(nickname string) bool {
	if i, err := app.Db().GetAll("select * from rat_user where user_name=?", nickname); err != nil {
		return false
	} else {
		return len(i) == 0
	}
}
