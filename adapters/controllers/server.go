package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/rinoguchi/microblog/usecases"
)

type Server struct {
	commentUsecase *usecases.CommentUsecase
}

func NewServer(commentUsecase *usecases.CommentUsecase) *Server {
	return &Server{
		commentUsecase: commentUsecase,
	}
}

func (s *Server) GetComments(w http.ResponseWriter, r *http.Request) {
	comments, err := s.commentUsecase.FindAllComment(r.Context())
	if err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func (s *Server) AddComment(w http.ResponseWriter, r *http.Request) {
	newComment := usecases.NewComment{Text: "dummy"}
	comment, err := s.commentUsecase.AddComment(r.Context(), &newComment)
	if err != nil {
		handleError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}
