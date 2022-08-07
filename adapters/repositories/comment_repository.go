package repositories

import (
	"context"
	"time"

	"github.com/rinoguchi/microblog/entities"
	"github.com/uptrace/bun"
)

type CommentRecord struct {
	bun.BaseModel `bun:"table:comments,alias:c"`

	Id        *int64     `bun:"id,pk"`
	Text      string     `bun:"text"`
	CreatedAt *time.Time `bun:"created_at"`
	UpdatedAt *time.Time `bun:"updated_at"`
}

func (cr CommentRecord) ToCommentEntity() entities.CommentEntity {
	return entities.CommentEntity{
		Id:        cr.Id,
		Text:      cr.Text,
		CreatedAt: cr.CreatedAt,
		UpdatedAt: cr.UpdatedAt,
	}
}

func NewCommentRecord(ce entities.CommentEntity) CommentRecord {
	return CommentRecord{
		Id:        ce.Id,
		Text:      ce.Text,
		CreatedAt: ce.CreatedAt,
		UpdatedAt: ce.UpdatedAt,
	}
}

type CommentRepositoryImpl struct {
}

func NewCommentRepositoryImpl() entities.CommentRepository {
	return CommentRepositoryImpl{}
}

func (c CommentRepositoryImpl) Add(ctx context.Context, commentEntity entities.CommentEntity) (entities.CommentEntity, error) {
	commentRecord := NewCommentRecord(commentEntity)
	_, err := GetDb().NewInsert().Model(&commentRecord).Exec(ctx)
	if err != nil {
		return entities.CommentEntity{}, err
	}
	return commentRecord.ToCommentEntity(), nil
}

func (c CommentRepositoryImpl) FindAll(ctx context.Context) ([]entities.CommentEntity, error) {
	var commentRecords []CommentRecord
	err := GetDb().NewSelect().Model(&commentRecords).Scan(ctx)
	if err != nil {
		return nil, err
	}

	commentEntities := make([]entities.CommentEntity, len(commentRecords))
	for i, commentRecord := range commentRecords {
		commentEntities[i] = commentRecord.ToCommentEntity()
	}
	return commentEntities, nil
}
