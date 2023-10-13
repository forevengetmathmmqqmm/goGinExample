package util

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/forevengetmathmmqqmm/goGinExample/pkg/e"
	"github.com/forevengetmathmmqqmm/goGinExample/pkg/setting"
)

var jwtSecret = []byte(setting.JwtSecret)
var Blacklist = make(map[string]bool)

type Claims struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(nickname, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		nickname,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "web3",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			if Blacklist[claims.Nickname] {
				return claims, nil
			} else {
				err := errors.New(e.GetMsg(e.ERROR_AUTH_CHECK_TOKEN_FAIL))
				return claims, err
			}
		}
	}

	return nil, err
}
