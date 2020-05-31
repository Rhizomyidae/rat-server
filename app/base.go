package app

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

type Controller struct {
	StartTime      int64
	HandlerSeconds float64
}

type Model struct {
	StartTime       string `json:"start_time,omitempty" gorm:"-" form:"start_time" time_format:"2008-08-08 08:08:08"` // 忽略这个字段
	EndTime         string `json:"end_time,omitempty" gorm:"-" form:"end_time" time_format:"2008-08-08 08:08:08"`     // 忽略这个字段
	Page            int64  `json:"page,omitempty" gorm:"-" form:"page"`                                               // 忽略这个字段
	PageSize        int64  `json:"page_size,omitempty" gorm:"-" form:"page_size"`                                     // 忽略这个字段
	OrderColumnName string `json:"order_column_name,omitempty" gorm:"-" form:"order_column_name"`                     // 忽略这个字段
	OrderType       string `json:"order_type,omitempty" gorm:"-" form:"order_type"`                                   // 忽略这个字段
	Fields          string `json:"fields,omitempty" gorm:"-" form:"fields"`                                           // 忽略这个字段
}

func init() {
	// 中间件注册
}

var db gdb.DB

func Db() gdb.DB {
	db = g.DB()
	return db
}

//
//func (c *Controller) Init(r *ghttp.Request) {
//	r.Response.Writeln("Init")
//}
//
//func (c *Controller) Shut(r *ghttp.Request) {
//	r.Response.Writeln("Shut")
//}
