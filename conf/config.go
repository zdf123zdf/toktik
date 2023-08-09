package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// 设置配置文件的搜索路径和文件名
func InitConfig() error {
	workDir, _ := os.Getwd()               // 找到工作目录
	viper.SetConfigName("config")          // 配置文件的文件名
	viper.SetConfigType("yml")             // 配置文件的后缀
	viper.AddConfigPath(workDir + "/conf") // 获取到配置文件的路径
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取配置文件失败：", err)
		return err
	}
	return nil
}
