package adapters

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rinoguchi/microblog/usecases"
)

type Server struct{}

func (s *Server) GetComments(w http.ResponseWriter, r *http.Request) {
	// TODO: get comments from DB via repository
	comments := []*usecases.Comment{newDummyComment()}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comments)
}

func (s *Server) AddComment(w http.ResponseWriter, r *http.Request) {
	// TODO: add comment to DB via repository
	comment := newDummyComment()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}

func NewServer() *Server {
	return &Server{}
}

func newDummyComment() *usecases.Comment {
	id := int64(123)
	now := time.Now()
	return &usecases.Comment{
		Id:        &id,
		Text:      "Dummy Text",
		CreatedAt: &now,
		UpdatedAt: &now,
	}
}
