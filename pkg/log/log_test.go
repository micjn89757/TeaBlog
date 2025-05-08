package log

import (
	"testing"

	"github.com/micjn89757/TeaBlog/internal/conf"
	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	conf := conf.NewConfig("config")
	log := NewLogger(conf)
	log.Info("test info", zap.String("url", "http:"))
	log.Error("test error")
	log.Debug("test debug")
}