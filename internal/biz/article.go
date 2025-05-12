package biz

import "context"

type Article struct {
}

type ArticleRepo interface {
	// db
	ListArticle(ctx context.Context)

	// reids
}
