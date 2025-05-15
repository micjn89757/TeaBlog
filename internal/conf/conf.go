package conf

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/micjn89757/TeaBlog/pkg/path"
	"github.com/spf13/viper"
)

type Config struct {
	Server Server `mapstructure:"server"`
	Data   Data   `mapstructure:"data"`
	Log    Log    `mapstructure:"log"`
}

type Log struct {
	Env string `mapstructure:"env"`
}

type Server struct {
	Http Http `mapstructure:"http"` // 注意字段名要大写，不然viper没有权限解析
}

type Http struct {
	Addr    string        `mapstructure:"addr"`
	Timeout time.Duration `mapstructure:"timeout"`
}

type Data struct {
	Postgresql PG    `mapstructure:"postgresql"`
	Redis      Redis `mapstructure:"redis"`
}

type PG struct {
	Driver string `mapstructure:"pgx"`
	Source string `mapstructure:"source"`
}

type Redis struct {
	Addr         string        `mapstructure:"addr"`
	DialTimeout  time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
	DB           int           `mapstructure:"db"`
}

func NewConfig(filename string) *Config {
	var config Config
	vp := viper.New()
	configPath := filepath.Join(path.GetRootPath(), "config")
	vp.AddConfigPath(configPath) // 文件所在目录
	vp.SetConfigName(filename)   // 文件名
	vp.SetConfigType("yaml")     // 扩展名

	configFile := configPath + filename + "yaml" // 配置文件完整路径

	if err := vp.ReadInConfig(); err != nil { // 查找并读取配置文件
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(fmt.Errorf("can not find the config file: %s", err)) // 系统初始化阶段发生任何错误，直接结束进程
		} else {
			panic(fmt.Errorf("parse the config file %s err: %s", configFile, err))
		}
	}

	vp.WatchConfig()                            // 监听配置文件变化
	vp.OnConfigChange(func(in fsnotify.Event) { // 配置文件发生变更之后会调用回调函数
		printStr := &strings.Builder{}
		printStr.WriteString("config file changed:")
		printStr.WriteString(in.Name)
		printStr.WriteByte('\n')
		io.WriteString(os.Stdout, printStr.String())

		// 重新解析
		if err := vp.Unmarshal(&config); err != nil {
			panic(fmt.Errorf("viper unmarshal err: %s", err))
		}
	})

	if err := vp.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("viper unmarshal err: %s", err))
	}

	return &config
}

func (c *Config) GetServerConfig() *Server {
	return &c.Server
}

func (c *Config) GetDataConfig() *Data {
	return &c.Data
}

//TODO: Validate
