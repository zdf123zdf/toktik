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
