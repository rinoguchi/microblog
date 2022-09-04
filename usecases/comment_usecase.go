package usecases

import (
	"context"

	"github.com/rinoguchi/microblog/entities"
	usecases "github.com/rinoguchi/microblog/usecases/models"
)

type CommentUsecase struct {
	commentRepository entities.CommentRepository
}

func NewCommentUsecase(
	commentRepository entities.CommentRepository,
) CommentUsecase {
	return CommentUsecase{
		commentRepository: commentRepository,
	}
}

func (c CommentUsecase) FindAllComment(ctx context.Context) ([]usecases.UComment, error) {
	commentEntities, err := c.commentRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	uComments := []usecases.UComment{}
	for _, commentEntity := range commentEntities {
		uComments = append(uComments, usecases.FromCommentEntity(commentEntity))
	}
	return uComments, nil
}

func (c CommentUsecase) AddComment(ctx context.Context, uComment usecases.UComment) (usecases.UComment, error) {
	commentEntity, err := c.commentRepository.Add(ctx, entities.CommentEntity{
		Text: uComment.Text,
	})
	if err != nil {
		return usecases.UComment{}, err
	}
	return toComment(commentEntity), nil
}

func toComment(commentEntity entities.CommentEntity) usecases.UComment {
	return usecases.UComment{
		Id:        commentEntity.Id,
		Text:      commentEntity.Text,
		CreatedAt: commentEntity.CreatedAt,
		UpdatedAt: commentEntity.UpdatedAt,
	}
}
