package entities

import (
	"context"
	"time"
)

type CommentEntity struct {
	Id        *int64
	Text      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type GetCommentsParamsEntity struct {
	Query     *string
	Year      *string
	Yearmonth *string
}

type CommentRepository interface {
	Find(ctx context.Context, paramsEntity GetCommentsParamsEntity) ([]CommentEntity, error)
	Add(ctx context.Context, comment CommentEntity) (CommentEntity, error)
}
