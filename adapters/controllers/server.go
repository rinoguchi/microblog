package controllers

import (
	"context"
	"encoding/json"
	"net/http"

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
	ctx := context.WithValue(r.Context(), "DB", repositories.GetDb())
	comments, err := s.commentUsecase.FindAllComment(ctx)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func (s *Server) AddComment(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "DB", repositories.GetDb())
	newComment := usecases.NewComment{Text: "dummy"}
	comment, err := s.commentUsecase.AddComment(ctx, newComment)
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}
