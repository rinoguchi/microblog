package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	controllers "github.com/rinoguchi/microblog/adapters/controllers/models"
	"github.com/rinoguchi/microblog/adapters/repositories"
	"github.com/rinoguchi/microblog/usecases"
	"github.com/uptrace/bun"
)

type Server struct {
	db             *bun.DB
	commentUsecase usecases.CommentUsecase
}

func NewServer(db *bun.DB, commentUsecase usecases.CommentUsecase) *Server {
	return &Server{
		db:             db,
		commentUsecase: commentUsecase,
	}
}

func (s *Server) GetComments(w http.ResponseWriter, r *http.Request) {
	uComments, err := s.commentUsecase.FindAllComment(r.Context())
	if err != nil {
		s.handleError(w, r, err)
		return
	}

	var comments []controllers.Comment
	for _, uComment := range uComments {
		comments = append(comments, controllers.FromUComment(uComment))
	}
	s.HandleOK(w, comments)
}

func (s *Server) AddComment(w http.ResponseWriter, r *http.Request) {
	var comment controllers.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		s.handleError(w, r, err)
		return
	}

	uComment, err := s.commentUsecase.AddComment(r.Context(), comment.ToUComment())
	if err != nil {
		s.handleError(w, r, err)
		return
	}
	s.HandleOK(w, controllers.FromUComment(uComment))
}

func (s *Server) SetTxMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.db.RunInTx(r.Context(), nil, func(ctx context.Context, tx bun.Tx) error {
			new_ctx := context.WithValue(r.Context(), repositories.TX_KEY, &tx)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r.WithContext(new_ctx))

			if ww.Status() != http.StatusOK {
				return errors.New("Rollbacked") // rollback
			}
			return nil // commit
		})
	})
}

func (s *Server) HandleOK(w http.ResponseWriter, obj interface{}) {
	s.setResponseHeaders(w)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(obj)
}

func (s *Server) handleError(w http.ResponseWriter, r *http.Request, err error) {
	s.setResponseHeaders(w)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}

func (s *Server) setResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}
