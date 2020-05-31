package util

import (
	"fmt"
	"github.com/Rhizomyidae/rat-server/common"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"time"
)

func CreateMagicToken(claims map[string]interface{}) string {
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(g.Cfg().GetInt("user.jwt.hours"))).Unix()
	claims["iat"] = time.Now().Unix()
	return createToken(claims)
}
func ParseMagicToken(tokenString string) (interface{}, bool) {
	return parseToken(tokenString)
}

func createToken(m map[string]interface{}) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	for index, val := range m {
		claims[index] = val
	}
	// fmt.Println(_map)
	token.Claims = claims
	tokenString, _ := token.SignedString(gconv.Bytes(g.Cfg().GetString("user.jwt.secret")))
	return tokenString
}

func parseToken(tokenString string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return gconv.Bytes(g.Cfg().GetString("user.jwt.secret")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}

func CheckToken(tokenString string) (b bool, t string, code int) {
	//kv := strings.Split(tokenString, " ")
	//if len(kv) != 2 || kv[0] != "Bearer" {
	//	fmt.Println("AuthString invalid:", tokenString)
	//	return false, nil
	//}

	token, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return gconv.Bytes(g.Cfg().GetString("user.jwt.secret")), nil
	})

	fmt.Println(err)

	if err != nil {
		switch err.(type) {

		case *jwt.ValidationError: // something was wrong during the validation
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				//ctx.Output.SetStatus(401)
				//resBody, err := json.Marshal(controllers.OutResponse(401, nil, "登录已过期，请重新登录"))
				//ctx.Output.Body(resBody)
				//if err != nil {
				//	panic(err)
				//}
				return false, "登录已过期，请重新登录", 50014
			default:
				//ctx.Output.SetStatus(401)
				//resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
				//ctx.Output.Body(resBytes)
				//if err != nil {
				//	panic(err)
				//}
				return false, "非法请求，请重新登录", 50008
			}
		default: // something else went wrong
			//ctx.Output.SetStatus(401)
			//resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
			//ctx.Output.Body(resBytes)
			//if err != nil {
			//	panic(err)
			//}
			return false, "非法请求，请重新登录", 50008
		}

		fmt.Println("转换为jwt claims失败.", err)
		return false, "非法请求，请重新登录", 50008
	}

	if !token.Valid {
		// but may still be invalid
		//ctx.Output.SetStatus(401)
		//resBytes, err := json.Marshal(controllers.OutResponse(401, nil, "非法请求，请重新登录"))
		//ctx.Output.Body(resBytes)
		//if err != nil {
		//	panic(err)
		//}
		return false, "非法请求，请重新登录", 50008
	}
	//GetUserNameByToken(kv[1])

	//redisClient,err := NewRedisClient()

	//username := GetUserNameByToken(tokenString)

	//jsonResToken,err := redisClient.Get("token_"+username).Result()
	//if err != nil || helpers.Empty(jsonResToken){
	//	return false,"非法请求，请重新登录",50008
	//}

	//verification := GetVerificationByToken(tokenString)

	//if(verification != helpers.Md5(username)){
	//	return false,"非法请求，请重新登录",50008
	//}

	userId := GetUserIdByToken(tokenString)
	user_name := GetUserNameByToken(tokenString)
	common.UserName = user_name
	return true, "验证通过", userId
}

func GetUserIdByToken(tokenString string) int {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return gconv.Bytes(g.Cfg().GetString("user.jwt.secret")), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	return int(id)
}

func GetUserNameByToken(tokenString string) string {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return gconv.Bytes(g.Cfg().GetString("user.jwt.secret")), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	user_name := claims["user_name"].(string)
	return user_name
}

func GetVerificationByToken(tokenString string) string {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return gconv.Bytes(g.Cfg().GetString("user.jwt.secret")), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	user_name := claims["verification"].(string)
	return user_name
}
