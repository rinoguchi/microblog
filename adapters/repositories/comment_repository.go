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
	tx := ctx.Value(TX_KEY).(*bun.Tx)
	_, err := tx.NewInsert().Model(&dbComment).Exec(ctx)
	if err != nil {
		return entities.CommentEntity{}, err
	}
	return dbComment.ToCommentEntity(), nil
}

func (c CommentRepositoryImpl) FindAll(ctx context.Context) ([]entities.CommentEntity, error) {
	var dbComments []repositories.DbComment
	tx := ctx.Value(TX_KEY).(*bun.Tx)
	err := tx.NewSelect().Model(&dbComments).Scan(ctx)
	if err != nil {
		return nil, err
	}

	commentEntities := make([]entities.CommentEntity, len(dbComments))
	for i, commentRecord := range dbComments {
		commentEntities[i] = commentRecord.ToCommentEntity()
	}
	return commentEntities, nil
}
