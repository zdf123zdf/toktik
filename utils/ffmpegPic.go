package utils

import (
	"github.com/google/uuid"
	"log"
	"os/exec"
)

// 通过FFmpeg生成视频封面
func GetFrame(videoFile, outputImage string) error {
	// 执行 FFmpeg 命令
	cmd := exec.Command("ffmpeg", "-i", videoFile, "-vframes", "1", "-q:v", "2", outputImage)
	if err := cmd.Run(); err != nil {
		log.Fatal("视频封面生成失败！", err)
		return err
	}
	return nil
}

// 生成随机名称
func RandomName() string {
	imageID := uuid.New().String()
	return imageID
}
