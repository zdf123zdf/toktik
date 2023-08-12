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
	IsFavorite    bool     `json:"is_favorite,omitempty"` // 是否点赞
	Title         string   `json:"title,omitempty"`
}

type FeedResponse struct {
	Response
	VideoList []FeedVideo `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`
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
	var response FeedResponse
	response.StatusCode = -1
	// 获取token和latest_time
	token := c.Query("token")
	fmt.Println("token", token)
	latestTimeStr := c.Query("latest_time")
	// 将时间戳转为时间格式
	var latestTime time.Time
	latestTime = time.Now()
	if latestTimeStr != "" {
		// 将时间戳字符串转换为整数
		timestamp, err := strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			response.StatusMsg = "时间戳格式错误"
			response.NextTime = latestTime.UnixMilli() // 毫秒时间戳
			c.JSON(http.StatusBadRequest, response)    // 400状态
			return
		}
		// 处理毫秒时间戳
		latestTime = time.Unix(0, timestamp*int64(time.Millisecond))
	}
	// 从service取数据
	videoList, err := service.FeedGet(latestTime)
	if err != nil {
		response.StatusMsg = "获取数据失败"
		response.NextTime = latestTime.UnixMilli()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.StatusCode = 0
	response.StatusMsg = "success"
	response.NextTime = latestTime.UnixMilli()

	feedVideos := make([]FeedVideo, 0)
	for _, video := range videoList {
		// 取出视频信息
		// 获取作者用户信息
		author := FeedUser{
			Id:              video.User.ID,
			Name:            video.User.Name,
			FollowCount:     video.User.FollowCount,
			FollowerCount:   video.User.FollowerCount,
			IsFollow:        true, // 未实现查询
			Avatar:          video.User.Avatar,
			BackgroundImage: video.User.BackgroundImage,
			Signature:       video.User.Signature,
			TotalFavorited:  video.User.TotalFavorited,
			WorkCount:       0, // 未实现查询
			FavoriteCount:   0, // 未实现查询
		}
		// 创建FeedVideo对象
		feedVideo := FeedVideo{
			Id:            video.ID,
			Author:        author,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    false, // 未实现查询
			Title:         video.Title,
		}

		// 将FeedVideo添加到feedVideos中
		feedVideos = append(feedVideos, feedVideo)
	}
	response.VideoList = feedVideos
	c.JSON(http.StatusOK, response)
}
