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

type CommentRepository interface {
	FindAll(ctx context.Context) ([]CommentEntity, error)
	Add(ctx context.Context, comment CommentEntity) (CommentEntity, error)
	Update(ctx context.Context, comment CommentEntity) (CommentEntity, error)
}
