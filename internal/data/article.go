package data

import (
	"go.uber.org/zap"
)

type articleRepo struct {
	data *Data
	log  *zap.Logger
}

func NewArticleRepo(data *Data, logger *zap.Logger) {
}
