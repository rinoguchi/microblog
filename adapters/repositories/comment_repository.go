package repositories

import (
	"context"

	repositories "github.com/rinoguchi/microblog/adapters/repositories/models"
	"github.com/rinoguchi/microblog/entities"
	"github.com/uptrace/bun"
)

type CommentRepositoryImpl struct {
}

func NewCommentRepositoryImpl() entities.CommentRepository {
	return CommentRepositoryImpl{}
}

func (c CommentRepositoryImpl) Add(ctx context.Context, commentEntity entities.CommentEntity) (entities.CommentEntity, error) {
	dbComment := repositories.FromCommentEntity(commentEntity)
	db := ctx.Value(DbKey).(*bun.DB)
	_, err := db.NewInsert().Model(&dbComment).Exec(ctx)
	if err != nil {
		return entities.CommentEntity{}, err
	}
	return dbComment.ToCommentEntity(), nil
}

func (c CommentRepositoryImpl) FindAll(ctx context.Context) ([]entities.CommentEntity, error) {
	var dbComments []repositories.DbComment
	db := ctx.Value(DbKey).(*bun.DB)
	err := db.NewSelect().Model(&dbComments).Scan(ctx)
	if err != nil {
		return nil, err
	}

	commentEntities := make([]entities.CommentEntity, len(dbComments))
	for i, commentRecord := range dbComments {
		commentEntities[i] = commentRecord.ToCommentEntity()
	}
	return commentEntities, nil
}
