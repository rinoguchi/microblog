package usecases

import (
	"context"

	"github.com/rinoguchi/microblog/entities"
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

func (c CommentUsecase) FindAllComment(ctx context.Context) ([]Comment, error) {
	commentEntities, err := c.commentRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	comments := []Comment{}
	for _, commentEntity := range commentEntities {
		comments = append(comments, toComment(commentEntity))
	}
	return comments, nil
}

func (c CommentUsecase) AddComment(ctx context.Context, newComment NewComment) (Comment, error) {
	commentEntity, err := c.commentRepository.Add(ctx, entities.CommentEntity{
		Text: newComment.Text,
	})
	if err != nil {
		return Comment{}, err
	}
	return toComment(commentEntity), nil
}

func toComment(commentEntity entities.CommentEntity) Comment {
	return Comment{
		Id:        commentEntity.Id,
		Text:      commentEntity.Text,
		CreatedAt: commentEntity.CreatedAt,
		UpdatedAt: commentEntity.UpdatedAt,
	}
}
