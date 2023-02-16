package core

import (
	"github.com/dgrijalva/jwt-go"
	"strings"
)

type MyCustomClaims struct {
	Domain   string `json:"domain"`
	Usid     int64  `json:"usid"`
	Username string `json:"username"`
	UserType int32  `json:"userType"`
	jwt.StandardClaims
}

type MyClaim struct {
	Domain   string `json:"domain"`
	Usid     int64  `json:"usid"`
	Username string `json:"username"`
	UserType int32  `json:"userType"`
}

func ParseJwt(tokenstr string) (bool, *MyClaim) {
	token, err := jwt.ParseWithClaims(strings.TrimSpace(tokenstr), &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if Cfg.ReleaseMode {
			return []byte(JWT_SECRET_STRING_PROD), nil
		} else {
			return []byte(JWT_SECRET_STRING_DEV), nil
		}

	})
	if err != nil {
		Error(err.Error())
		return false, nil
	}
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		mc := new(MyClaim)
		mc.Domain = claims.Domain
		mc.Usid = claims.Usid
		mc.Username = claims.Username
		mc.UserType = claims.UserType
		return true, mc
	}
	//Info(fmt.Sprintf("%v %v", claims.Usid, claims.StandardClaims.ExpiresAt))
	return false, nil
}
