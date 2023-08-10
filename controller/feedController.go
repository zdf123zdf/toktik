package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"toktik/service"
)

type FeedVideo struct {
	Id            uint     `json:"id,omitempty"`
	Author        FeedUser `json:"author,omitempty"`
	PlayUrl       string   `json:"play_url,omitempty"`
	CoverUrl      string   `json:"cover_url,omitempty"`
	FavoriteCount uint     `json:"favorite_count,omitempty"`
	CommentCount  uint     `json:"comment_count,omitempty"`
	IsFavorite    bool     `json:"is_favorite,omitempty"`
	Title         string   `json:"title,omitempty"`
}

type FeedResponse struct {
	Response
	VideoList []FeedVideo `json:"video_list,omitempty"`
	NextTime  uint        `json:"next_time,omitempty"`
}

type FeedUser struct {
	Id              uint   `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	FollowCount     uint   `json:"follow_count,omitempty"`     // 关注总数
	FollowerCount   uint   `json:"follower_count,omitempty"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow,omitempty"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar,omitempty"`           // 用户头像
	BackgroundImage string `json:"background_image,omitempty"` // 用户个人页顶部大图
	Signature       string `json:"signature,omitempty"`        // 个人简介
	TotalFavorited  uint   `json:"total_favorited"`            // 获赞数量
	WorkCount       uint   `json:"work_count_favorited"`       // 作品数
	FavoriteCount   uint   `json:"favorite_count"`             // 喜欢数
}

func GetFeed(c *gin.Context) {
	// 获取token和latest_time
	//token := c.Query("token")
	latestTimeStr := c.Query("latest_time")
	// 将时间戳转为时间格式
	var latestTime time.Time
	if latestTimeStr != "" {
		// 将时间戳字符串转换为整数
		timestamp, err := strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			fmt.Println("时间戳格式错误:", err)
			return
		}

		// 处理毫秒时间戳
		latestTime = time.Unix(0, timestamp*int64(time.Millisecond))

	} else {
		latestTime = time.Now()
	}
	// 从service取数据
	videoList, err := service.FeedGet(latestTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	fmt.Println(videoList)
	// 示例数据
	response := FeedResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "Success",
		},
		VideoList: []FeedVideo{
			{
				Id:       1,
				Author:   FeedUser{Id: 1, Name: "John Doe"},
				PlayUrl:  "https://example.com/video/1",
				CoverUrl: "https://example.com/cover/1",
			},
			{
				Id:       2,
				Author:   FeedUser{Id: 2, Name: "Jane Smith"},
				PlayUrl:  "https://example.com/video/2",
				CoverUrl: "https://example.com/cover/2",
			},
		},
		NextTime: 1691644322647,
	}
	c.JSON(http.StatusOK, response)

}
