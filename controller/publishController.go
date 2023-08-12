package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/h2non/filetype"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"toktik/model"
	"toktik/service"
	"toktik/utils"
)

// 视频投稿接口
func PublishAction(c *gin.Context) {
	var response Response
	response.StatusCode = -1
	// 获取用户token
	token := c.PostForm("token")
	fmt.Println(token)
	// 获取视频
	title := c.PostForm("title")
	if title == "" {
		response.StatusMsg = "请输入标题"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	file, header, err := c.Request.FormFile("data")
	if err != nil {
		response.StatusMsg = "没有视频文件上传"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// 校验是否为视频格式
	content, err := io.ReadAll(file)
	if err != nil {
		response.StatusMsg = "无法读取视频文件"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if filetype.IsVideo(content) == false {
		response.StatusMsg = "上传的文件不是视频类型"
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// 视频名称生成 用户名+随机字符串
	// 1 根据token获取用户名
	// 2 添加随机字符串拼接视频名称
	ext := filepath.Ext(header.Filename) // 获取文件扩展名
	name := utils.RandomName()
	videoName := name + ext
	// 视频保存到本地
	_, err = file.Seek(0, io.SeekStart) // 将文件指针重置到文件开头
	out, err := os.Create("public/" + videoName)
	if err != nil {
		response.StatusMsg = "视频无法上传"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		response.StatusMsg = "视频无法保存"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 生成视频封面
	picName := name + ".png"
	err = utils.GetFrame("public/"+videoName, "public/"+picName)
	if err != nil {
		response.StatusMsg = "视频上传失败"
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// 删除视频和封面
	defer func() {
		out.Close() // 关闭视频文件
		err := os.Remove("public/" + videoName)
		if err != nil {
			log.Println("无法删除视频文件:", err.Error())
		}
		err = os.Remove("public/" + picName)
		if err != nil {
			log.Println("无法删除封面图片文件:", err.Error())
		}
	}()

	// 上传到对象存储中
	payUrl, err := service.UploadObject(videoName, "public/"+videoName)
	if err != nil {
		response.StatusMsg = "视频上传失败"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	coverUrl, err := service.UploadObject(picName, "public/"+picName)
	if err != nil {
		response.StatusMsg = "视频上传失败"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	// 保存链接到数据库
	video := model.Video{
		PlayUrl:       payUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
		UserID:        1, // 视频作者ID 目前定义死，后面修改
	}
	err = service.CreateVideo(video)
	if err != nil {
		response.StatusMsg = "视频上传失败"
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.StatusCode = 0
	response.StatusMsg = "success"
	c.JSON(http.StatusOK, response)
}
