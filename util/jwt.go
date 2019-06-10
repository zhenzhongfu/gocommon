package util

import (
	"time"

	"github.com/zhenzhongfu/gocommon/setting"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	UserID       string `json:"userID"`
	IdentityType string `json:"identityType"`
	Identifier   string `json:"identifier"`
	Credential   string `json:"credential"`
	jwt.StandardClaims
}

func GenerateToken(userID, identityType, identifier, credential string) (string, error) {
	// 24小时有效?
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	//expireTime := nowTime.Add(30 * time.Second)
	claims := Claims{
		userID,
		identityType,
		identifier,
		credential, // 已经过bcrypt之后的值
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "go-backend",
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
			return claims, nil
		}
	}
	return nil, err
}
