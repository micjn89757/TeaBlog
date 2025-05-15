package biz

import "time"

type Base struct {
	ID        string    `gorm:"column:id;primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli"`
	DeleteAt  time.Time `gorm:"index"`
}
