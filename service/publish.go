package service

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"toktik/dao/db"
	"toktik/dao/minioStore"
	"toktik/model"
)

// 上传文件
func UploadObject(objectName, filePath string) (objectURL string, err error) {
	minioUrl := viper.GetString("minio.url")
	minioPort := viper.GetString("minio.port")
	bucketName := viper.GetString("minio.bucketName")
	_, err = minioStore.MinioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln("文件上传失败:", err.Error())
	}
	// 返回文件链接
	objectURL = fmt.Sprintf("http://%s:%s/%s/%s", minioUrl, minioPort, bucketName, url.PathEscape(objectName))
	return objectURL, nil
}

// 保存链接到数据库中
func CreateVideo(video model.Video) error {
	result := db.DB.Create(&video)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 获取用户发布视频列表
func GetPublish(userID uint) (publishList []*model.Video, err error) {
	// 查询用户的所有视频记录
	if err = db.DB.Preload("User").Where("user_id = ?", userID).Find(&publishList).Error; err != nil {
		return nil, err
	}
	return
}
