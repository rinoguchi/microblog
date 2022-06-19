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
) *CommentUsecase {
	return &CommentUsecase{
		commentRepository: commentRepository,
	}
}

func (p *CommentUsecase) AddComment(ctx context.Context, newComment *NewComment) (*Comment, error) {
	commentEntity, err := p.commentRepository.Add(ctx, &entities.CommentEntity{
		Text: newComment.Text,
	})
	if err != nil {
		return nil, err
	}
	return toComment(commentEntity), nil
}

func (p *CommentUsecase) FindAllComment(ctx context.Context) ([]*Comment, error) {
	commentEntities, err := p.commentRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	comments := []*Comment{}
	for _, commentEntity := range commentEntities {
		comments = append(comments, toComment(commentEntity))
	}
	return comments, nil
}

func toComment(commentEntity *entities.CommentEntity) *Comment {
	return &Comment{
		Id:        commentEntity.Id,
		Text:      commentEntity.Text,
		CreatedAt: commentEntity.CreatedAt,
		UpdatedAt: commentEntity.UpdatedAt,
	}
}
