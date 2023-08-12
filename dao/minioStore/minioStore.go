package minioStore

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"log"
)

var MinioClient *minio.Client // 定义全局的 MinIO 客户端对象
// 初始化minio 创建存储桶
func InitMinio() (err error) {
	// 读取配置文件
	minioUrl := viper.GetString("minio.url")
	minioPort := viper.GetString("minio.port")
	minioAccessKey := viper.GetString("minio.accessKey")
	minioSecretKey := viper.GetString("minio.secretKey")
	bucketName := viper.GetString("minio.bucketName")
	location := viper.GetString("minio.location")

	//根据配置文件连接minio
	MinioClient, err = minio.New(minioUrl+":"+minioPort, &minio.Options{Creds: credentials.NewStaticV4(minioAccessKey, minioSecretKey, "")})
	if err != nil {
		log.Fatalln("minio服务连接失败:", err.Error())
	}
	// 判断存储桶是否存在
	var b bool
	b, err = MinioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalln("存储桶判断是否存在失败:", err.Error())
	}
	if b {
		return nil
	} else {
		// 存储桶不存在
		err = MinioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			log.Fatalln("存储桶创建失败:", err)
		}
		log.Println("存储桶创建成功！")
	}
	log.Println("minio服务初始化完成！")
	return nil
}
