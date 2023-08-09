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
	err = db.DB.Order("created_at desc").Where("created_at <= ?", latestTime).Limit(30).Find(&articleList).Error
	return
}
