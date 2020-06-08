// Package router is http-router package
package router

import (
	"github.com/go-chi/chi"

	"github.com/oshou/AwesomeMusic-api/ui/http/handler"
)

// IRouter is http-router interface
type IRouter interface{}

type router *chi.Router

// NewRouter is constructor for router
func NewRouter(h handler.IAppHandler) chi.Router {
	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", h.GetUsers)
			r.Post("/", h.AddUser)
			r.Get("/{user_id}", h.GetUserByID)
		})
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", h.GetPosts)
			r.Post("/", h.AddPost)
			r.Route("/{post_id}", func(r chi.Router) {
				r.Get("/", h.GetPostByID)
				r.Delete("/", h.DeletePostByID)
				r.Get("/comments", h.GetComments)
				r.Post("/comments", h.AddComment)
			})
		})
		r.Get("/tags/{tag_id}/posts", h.GetPostsByTagID)
	})
	return r
}
