package conf

import (
	"fmt"
	"path/filepath"

	"github.com/micjn89757/TeaBlog/pkg/util"
	"github.com/spf13/viper"
)

// TODO: 配置可以进行校验

func NewConfig(file string) *viper.Viper {
	config := viper.New()
	configPath := filepath.Join(util.GetRootPath(), "config")
	config.AddConfigPath(configPath) 	// 文件所在目录
	config.SetConfigName(file) 			// 文件名
	config.SetConfigType("yaml") 		// 扩展名
	
	configFile := configPath + file + "yaml"  // 配置文件完整路径

	if err := config.ReadInConfig(); err != nil { // 查找并读取配置文件
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("can not find the config file: %s", err)) // 系统初始化阶段发生任何错误，直接结束进程
		} else {
			panic(fmt.Errorf("parse the config file %s err: %s", configFile, err))
		}
	}

	return config
}

//TODO: Validate 