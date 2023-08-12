package service

import (
	"time"
	"toktik/dao/db"
	"toktik/model"
)

// 获取视频列表，根据传入时间戳返回视频，限制30条
func FeedGet(latestTime time.Time) (articleList []*model.Video, err error) {
	// 初始化 articleList 为一个空的切片
	articleList = make([]*model.Video, 0)
	// 使用Preload("User")来预加载用户信息
	err = db.DB.Preload("User").Order("created_at desc").Where("created_at <= ?", latestTime).Limit(30).Find(&articleList).Error
	return
}
