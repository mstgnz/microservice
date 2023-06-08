package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mstgnz/services/config"
	"github.com/mstgnz/services/handler"
	customMiddleware "github.com/mstgnz/services/middleware"
	"github.com/mstgnz/services/repository"
	"github.com/mstgnz/services/service"
)

var (
	db                = config.OpenDatabase()
	blogRepository    = repository.BlogRepository(db)
	commentRepository = repository.CommentRepository(db)
	blogService       = service.BlogService(blogRepository)
	commentService    = service.CommentService(commentRepository)

	blogHandler    = handler.BlogHandler(blogService)
	commentHandler = handler.CommentHandler(commentService)
)

type Config struct {
}

func main() {
	defer config.CloseDatabase(db)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "HEAD", "OPTION"},
		AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Route("/api", func(r chi.Router) {
		r.Route("/blogs", func(r chi.Router) {
			r.Get("/", blogHandler.All)
			r.Get("/{id}", blogHandler.FindByID)
			r.With(customMiddleware.TokenValidate).Post("/", blogHandler.Insert)
			r.With(customMiddleware.TokenValidate).Put("/{id}", blogHandler.Update)
			r.With(customMiddleware.TokenValidate).Delete("/{id}", blogHandler.Delete)
		})
		r.Route("/comments", func(r chi.Router) {
			r.Use(customMiddleware.TokenValidate)
			r.Post("/", commentHandler.Insert)
			r.Put("/{id}", commentHandler.Update)
			r.Delete("/{id}", commentHandler.Delete)
		})
	})

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Printf(err.Error())
	}
}
