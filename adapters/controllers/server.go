package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	controllers "github.com/rinoguchi/microblog/adapters/controllers/models"
	"github.com/rinoguchi/microblog/adapters/repositories"
	"github.com/rinoguchi/microblog/usecases"
)

type Server struct {
	commentUsecase usecases.CommentUsecase
}

func NewServer(commentUsecase usecases.CommentUsecase) *Server {
	return &Server{
		commentUsecase: commentUsecase,
	}
}

func (s *Server) GetComments(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), repositories.DbKey, repositories.GetDb())
	uComments, err := s.commentUsecase.FindAllComment(ctx)
	if err != nil {
		handleError(w, err)
		return
	}

	var comments []controllers.Comment
	for _, uComment := range uComments {
		comments = append(comments, controllers.FromUComment(uComment))
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func (s *Server) AddComment(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), repositories.DbKey, repositories.GetDb())

	var comment controllers.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		handleError(w, err)
		return
	}

	uComment, err := s.commentUsecase.AddComment(ctx, comment.ToUComment())
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(controllers.FromUComment(uComment))
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}
