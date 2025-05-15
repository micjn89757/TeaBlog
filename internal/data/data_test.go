package data

import (
	"testing"

	"github.com/micjn89757/TeaBlog/internal/conf"
	"github.com/micjn89757/TeaBlog/pkg/log"
)

func TestData(t *testing.T) {
	conf := conf.NewConfig("config.yaml")
	t.Log(conf)
	logger := log.NewLogger(conf)
	data, f, err := NewData(conf, logger)
	if err != nil {
		t.Error(err)
	}
	t.Log(data.db.Statement.Config)
	f()
}
