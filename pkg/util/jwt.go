package util

import (
	"time"
	"unbound/pkg/setting"

	jwt "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.RegisteredClaims
}

func GenerateToken(username, password string) (string, error) {
	// 获取当前时间
	nowTime := time.Now()
	// 设置过期时间为当前时间加上3小时
	expireTime := nowTime.Add(3 + time.Hour)

	// 创建Claims结构体，包含用户名、密码和注册声明
	claims := Claims{
		username,
		password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // 设置过期时间
			Issuer:    "gin-blog",                     // 设置发行者
		},
	}

	// 使用ES256签名算法创建tokenClaims，包含Claims信息
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	// 返回tokenClaims的签名字符串
	return tokenClaims.SignedString(jwtSecret)
}

func ParseToken(token string) (*Claims, error) {
	// 解析token，并传入Claims结构体指针作为参数  tokenClaims包含一个Claims字段传递结构体 是为了让他填充到 tokenClaims.Claims中
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		// 返回秘钥，用于验证token的签名
		return jwtSecret, nil
	})

	// 如果tokenClaims不为空
	if tokenClaims != nil {
		// 将tokenClaims的Claims字段断言为*Claims类型
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			// 如果断言成功且tokenClaims有效，则返回claims和nil错误
			return claims, nil
		}
	}

	// 如果tokenClaims为空或者无效，则返回nil和错误err
	return nil, err
}
