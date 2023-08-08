package controller

import (
	"errors"
	"net/http"
	"time"

	"github.com/abuziming/dousheng_demo/config"
	"github.com/abuziming/dousheng_demo/dao"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte(config.Global.Secret.JwtKey)

type Claim struct {
	UserId int64
	jwt.StandardClaims
}

// 颁发 token
func GetToken(user *dao.UserLogin) (string, error) {
	claim := &Claim{
		UserId: user.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * 24 * time.Hour).UnixMilli(), // 设置过期时间为 30 天
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析 token
func ParseToken(tokenString string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claim); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token 不正确")
}

// 鉴权并设置 user_id
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}

		if token == "" {
			c.JSON(http.StatusOK, Response{
				StatusCode: 401,
				StatusMsg:  "token 未发送",
			})
			c.Abort()
			return
		}

		claim, err := ParseToken(token)
		if err != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 403,
				StatusMsg:  err.Error(),
			})
			c.Abort()
			return
		}
		// token 超时
		if time.Now().UnixMilli() > claim.ExpiresAt {
			c.JSON(http.StatusOK, Response{
				StatusCode: 401,
				StatusMsg:  "token 已过期",
			})
			c.Abort()
			return
		}

		c.Set("user_id", claim.UserId)
		c.Next()
	}
}
