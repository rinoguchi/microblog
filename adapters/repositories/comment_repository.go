package repositories

import (
	"context"
	"math/rand"
	"time"

	"github.com/rinoguchi/microblog/entities"
)

type CommentRepositoryImpl struct {
}

func NewCommentRepositoryImpl() entities.CommentRepository {
	return &CommentRepositoryImpl{}
}

func (c *CommentRepositoryImpl) Add(ctx context.Context, comment *entities.CommentEntity) (*entities.CommentEntity, error) {
	// TODO: 実際のDBにアクセスする
	return newDummyCommentEntity(), nil
}

func (c *CommentRepositoryImpl) Update(ctx context.Context, comment *entities.CommentEntity) (*entities.CommentEntity, error) {
	// TODO: 実際のDBにアクセスする
	return newDummyCommentEntity(), nil
}

func (c *CommentRepositoryImpl) FindAll(ctx context.Context) ([]*entities.CommentEntity, error) {
	// TODO: 実際のDBにアクセスする
	return []*entities.CommentEntity{newDummyCommentEntity(), newDummyCommentEntity()}, nil
}

func newDummyCommentEntity() *entities.CommentEntity {
	id := int64(rand.Intn(100))
	now := time.Now()
	return &entities.CommentEntity{
		Id:        &id,
		Text:      "Dummy Text",
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
