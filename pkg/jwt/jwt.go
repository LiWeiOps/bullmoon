package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中

type MyClaims struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.RegisteredClaims
}

var (
	// MySecret用于加盐的字符串
	MySecret            = []byte("夏天夏天悄悄过去")
	TokenExpireDuration = time.Hour * 24
)

func GenToken(userid int64, username string) (string, error) {
	mc := MyClaims{
		userid,
		username,
		jwt.RegisteredClaims{
			Issuer:    "bull_moon",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
		},
	}
	// 使用指定的签名方法创建token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mc)
	return token.SignedString(MySecret)
}

func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	claims, ok := token.Claims.(*MyClaims)
	if ok {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
