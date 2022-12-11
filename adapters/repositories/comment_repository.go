package repositories

import (
	"context"
	"regexp"

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

func (c CommentRepositoryImpl) Find(ctx context.Context, paramsEntity entities.GetCommentsParamsEntity) ([]entities.CommentEntity, error) {
	var dbComments []repositories.DbComment
	tx := ctx.Value(TX_KEY).(*bun.Tx)
	var err error
	if paramsEntity.Query != nil && len(regexp.MustCompile("[ 　]+").ReplaceAllString(*paramsEntity.Query, "")) > 0 {
		selectQuery := tx.NewSelect().Model(&dbComments)
		for i, q := range regexp.MustCompile(`[ 　]+`).Split(*paramsEntity.Query, -1) {
			if i == 0 {
				selectQuery.Where("text ~* ?", q)
			} else {
				selectQuery.WhereOr("text ~* ?", q)
			}
		}
		err = selectQuery.Scan(ctx)
	} else if paramsEntity.Year != nil {
		err = tx.NewSelect().Model(&dbComments).Where("to_char(c.created_at, 'YYYY') = ?", paramsEntity.Year).Scan(ctx)
	} else if paramsEntity.Yearmonth != nil {
		err = tx.NewSelect().Model(&dbComments).Where("to_char(c.created_at, 'YYYYMM') = ?", paramsEntity.Yearmonth).Scan(ctx)
	} else {
		err = tx.NewSelect().Model(&dbComments).Scan(ctx)
	}
	if err != nil {
		return nil, err
	}

	commentEntities := make([]entities.CommentEntity, len(dbComments))
	for i, commentRecord := range dbComments {
		commentEntities[i] = commentRecord.ToCommentEntity()
	}
	return commentEntities, nil
}
