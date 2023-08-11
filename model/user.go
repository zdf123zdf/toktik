package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name            string `gorm:"column:name;unique" json:"name"`                  // 用户名称
	Password        string `gorm:"column:password" json:"password"`                 // 用户密码
	FollowCount     uint   `gorm:"column:follow_count" json:"follow_count"`         // 关注总数
	FollowerCount   uint   `gorm:"column:follower_count" json:"follower_count"`     // 粉丝总数
	Avatar          string `gorm:"column:avatar" json:"avatar"`                     // 用户头像链接
	BackgroundImage string `gorm:"column:background_image" json:"background_image"` // 用户个人页顶部大图
	Signature       string `gorm:"column:signature" json:"signature"`               // 个人简介
	TotalFavorited  uint   `gorm:"column:total_favorited" json:"total_favorited"`   // 获赞数量
	FavoriteCount   uint   `gorm:"column:favorite_count" json:"favorite_count"`     // 点赞数量
}
