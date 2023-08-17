package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavoriteResponse struct {
	Response
}

func FavoriteAction(c *gin.Context) {
	// 获取信息
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	userId := c.GetUint("user_id")

	var response = FavoriteResponse{}
	if actionType == "1" {
		// 点赞操作

		c.JSON(http.StatusOK, FavoriteResponse{
			Response: Response{StatusCode: 0},
		})
	} else if actionType == "2" {
		// 取消点赞操作

		c.JSON(http.StatusOK, FavoriteResponse{
			Response: Response{StatusCode: 0},
		})
	} else {
		c.JSON(http.StatusOK, FavoriteResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Failed to operation"},
		})
	}
}

func FavoriteList(c *gin.Context) {
	userId = c.Query("user_id")
	token := c.Query("token")
}
