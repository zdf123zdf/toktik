package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"toktik/dao/redis"
)

type Response struct {
	StatusCode int32  `json:"status_code"`          // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg,omitempty"` // 返回状态描述
}

type MyClaims struct {
	UserId   uint   `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

var Key = []byte("ganfandouyin")

// 校验token
func verifyToken(token string) (*MyClaims, bool) {
	tokenObj, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	})
	if key, _ := tokenObj.Claims.(*MyClaims); tokenObj.Valid {
		return key, true
	} else {
		return nil, false
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		// token不存在，用户未登录
		if tokenStr == "" {
			c.JSON(http.StatusOK, Response{StatusCode: 401, StatusMsg: "用户不存在，请登录"})
			c.Abort() //阻止执行
			return
		}
		// 校验token
		tokenStruck, ok := verifyToken(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, Response{StatusCode: 401, StatusMsg: "token错误，请重新登录"})
			c.Abort() //阻止执行
			return
		}
		// 利用用户ID取出redis中的token 判断是否正确和过期
		token, err := redis.Redisdb.Get(strconv.Itoa(int(tokenStruck.UserId))).Result()
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 401, StatusMsg: "后台错误"})
			c.Abort() //阻止执行
			return
		}
		if tokenStr != token {
			c.JSON(http.StatusOK, Response{StatusCode: 401, StatusMsg: "token过期，请重新登录"})
			c.Abort() //阻止执行
			return
		}
		// 保存到上下文
		c.Set("user_id", tokenStruck.UserId)
		c.Set("username", tokenStruck.UserName)
		// Token有效，继续处理后续请求
		c.Next()
	}
}
