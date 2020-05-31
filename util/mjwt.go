package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"time"
)

func CreateMagicToken(claims map[string]interface{}) string {
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(gconv.Int(g.Cfg("user.jwt.hours")))).Unix()
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
	tokenString, _ := token.SignedString(gconv.Bytes(g.Cfg("user.jwt.secret")))
	return tokenString
}

func parseToken(tokenString string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return gconv.Bytes(g.Cfg("user.jwt.secret")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err)
		return "", false
	}
}
