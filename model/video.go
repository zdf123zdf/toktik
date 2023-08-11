package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	PlayUrl       string `gorm:"column:play_url" json:"play_url"`             // 视频播放地址
	CoverUrl      string `gorm:"column:cover_url" json:"cover_url"`           // 视频封面地址
	FavoriteCount uint   `gorm:"column:favorite_count" json:"favorite_count"` // 视频的点赞总数
	CommentCount  uint   `gorm:"column:comment_count" json:"comment_count"`   // 视频的评论总数
	Title         string `gorm:"column:title" json:"title"`                   // 视频标题
	UserID        uint   // 视频作者信息 外键
	User          User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` // 联级删除
}
