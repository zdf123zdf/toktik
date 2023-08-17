package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toktik/model"
)

type Comment struct {
	id        int64 `json:"id"`
	user      model.User
	content   string `json:"content"`
	creatData string `json:"creat_data"`
}

type CommentResponse struct {
	Response
	comment Comment
}

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")

	var response = CommentResponse{}
	if actionType == "1" {
		// 发布评论
		CommentText := c.Query("comment_text") // 用户填写的评论内容
		response.comment.content = CommentText
		// 获取当前时间赋值给response.comment.creatData

		// 评论内容和时间等信息(response)存储到数据库表中

		// 返回响应
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{StatusCode: 0},
			comment: Comment{
				content: CommentText,
			},
		})
	} else if actionType == "2" {
		// 删除操作
		CommentId := c.Query("comment_id") // 要删除的评论ID
		// 数据库表中删除该评论内容

		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Comment deleted successfully"},
		})
	} else {
		c.JSON(http.StatusOK, CommentResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Failed to operation"},
		})
	}
}

func CommentList(c *gin.Context) {
	user_id = c.Query("user_id")
	token := c.Query("token")
}
