package data

import (
	"github.com/google/wire"
	"github.com/micjn89757/TeaBlog/internal/conf"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewArticleRepo)

type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// create data obj, return func() for close connection
func NewData(conf *conf.Config, logger *zap.Logger) (*Data, func(), error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.Data.Redis.Addr,
		Password:     "",
		DB:           conf.Data.Redis.DB,
		WriteTimeout: conf.Data.Redis.WriteTimeout,
		ReadTimeout:  conf.Data.Redis.ReadTimeout,
		DialTimeout:  conf.Data.Redis.DialTimeout,
	})

	pgdb, err := gorm.Open(postgres.New(postgres.Config{
		DSN: conf.Data.Postgresql.Source,
		// DriverName: conf.Data.Postgresql.Driver,
		// PreferSimpleProtocol: true, // disable implicit prepared statement
	}), &gorm.Config{
		SkipDefaultTransaction: true, // disable implicit transaction
	})

	if err != nil {
		return nil, nil, err
	}

	// config connect pool
	sqlDB, err := pgdb.DB()
	if err != nil {
		return nil, nil, err
	}

	//  set params of pool
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)

	data := &Data{
		rdb: rdb,
		db:  pgdb,
	}

	return data, func() {
		logger.Info("closing the data resource")
		if err := data.rdb.Close(); err != nil {
			logger.Error("close redis err", zap.Error(err))
		}

	}, nil
}
