package response

import (
	"github.com/gogf/gf/net/ghttp"
)

// 错误码
type ErrorCode int

const (
	SUCCESS     ErrorCode = 200
	FALL        ErrorCode = 400
	NOT_LOGIN   ErrorCode = 411
	VERIFY_FALL ErrorCode = 412
)

const (
	SUCCESS_MSG = "OK"
)

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    ErrorCode   `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code ErrorCode, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	_ = r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err ErrorCode, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.Exit()
}

func SuccessResult(data interface{}) JsonResponse {
	return JsonResponse{SUCCESS, SUCCESS_MSG, data}
}

func ErrorResult(code ErrorCode, message string) JsonResponse {
	return JsonResponse{code, message, nil}
}
