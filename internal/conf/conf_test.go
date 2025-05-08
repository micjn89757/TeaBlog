package conf

import (
	"testing"
	"time"
)

func TestNewConfig(t *testing.T) {
	conf := NewConfig("config.yaml")

	t.Log(conf)
	time.Sleep(10 * time.Second) // 10s内修改一下配置文件, 测试监听功能是否正常
	t.Log(conf)


	t.Log(conf.GetDataConfig())
	t.Log(conf.GetServerConfig())
}