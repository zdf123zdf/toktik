package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"toktik/dao/db"
	"toktik/model"
)

type UserInfos struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

func UserRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// hash加密密码
	hashpassword, err := HashPassword(password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Error encryption password"},
		})
		return
	}

	// 计算token
	token, err := GenerateToken(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 如果name已经存在，则返回StatusCode: 1(失败)和"User already exist"
	var user model.User
	if err := db.DB.Where("name = ?", username).First(&user).Error; err == nil {
		// 用户已经存在
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
		return
	}

	// 创建新用户
	newUser := model.User{
		Name:     username,
		Password: hashpassword,
	}

	// 使用 GORM 的 Create 方法来创建用户
	if err := db.DB.Create(&newUser).Error; err != nil {
		// 处理错误
		c.JSON(http.StatusInternalServerError, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Failed to create user"},
		})
		return
	}
	// 存储token

	// 返回响应
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   int64(user.ID),
		Token:    token,
	})
}

func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// 计算token
	token, err := GenerateToken(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var user model.User
	if err := db.DB.Where("name = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		// 检验密码是否正确
		storedPassword := user.Password
		isMatch := CheckPassword(password, storedPassword)
		if isMatch {
			// 存储token

			// 返回响应
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   int64(user.ID),
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Wrong password"},
			})
		}
	}
}

func UserInfo(c *gin.Context) {
	username := c.Query("username")
	var user model.User
	if err := db.DB.Where("name = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	} else {
		var userinfo = UserInfos{
			Id:            int64(user.ID),
			Name:          username,
			FollowCount:   0,
			FollowerCount: 0,
		}
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     userinfo,
		})
	}
}
