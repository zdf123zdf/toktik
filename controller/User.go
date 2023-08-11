package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
)

/*
问题：
	1.用户ID每次重启都会刷新(从1开始累加)，原因是最初的usersLoginInfo,重启后会被重置
		可能需要其他的存储方法来解决，或许加密算法也可以优化？
	2.用户登录接口token不存在是会返回"User doesn't exist"，这似乎不太严谨，因为还有可能是密码错误
*/

// 用户logininfo使用map存储用户信息，而键是演示的用户名密码,即生产的token
// 每次服务器启动
// 测试数据时，都会清除用户数据：username = toktik，password = 123456
var usersLoginInfo = map[string]User{
	"toktik123456": {
		Id:            1,
		Name:          "toktik",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

func UserRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	// 如果token已经存在，则返回StatusCode: 1(失败)和"User already exist"
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}

func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
