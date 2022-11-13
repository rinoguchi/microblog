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

func (c CommentUsecase) FindComments(ctx context.Context, params usecases.UGetCommentsParams) ([]usecases.UComment, error) {
	commentEntities, err := c.commentRepository.Find(ctx, params.ToGetCommentsParamsEntity())
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
	return usecases.FromCommentEntity(commentEntity), nil
}
