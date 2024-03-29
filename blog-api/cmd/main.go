package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mstgnz/microservice/config"
	"github.com/mstgnz/microservice/handler"
	customMiddleware "github.com/mstgnz/microservice/middleware"
	"github.com/mstgnz/microservice/repository"
	"github.com/mstgnz/microservice/service"
)

var (
	db                = config.OpenDatabase()
	blogRepository    = repository.NewBlogRepository(db)
	commentRepository = repository.NewCommentRepository(db)
	blogService       = service.NewBlogService(blogRepository)
	commentService    = service.NewCommentService(commentRepository)

	blogHandler    = handler.NewBlogHandler(blogService)
	commentHandler = handler.NewCommentHandler(commentService)
)

func main() {
	defer config.CloseDatabase(db)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.Heartbeat("/ping"))

	r.Route("/api", func(r chi.Router) {
		r.Route("/blogs", func(r chi.Router) {
			r.Get("/", blogHandler.All)
			r.Get("/{slug}", blogHandler.Find)
			r.With(customMiddleware.TokenValidate).Post("/", blogHandler.Create)
			r.With(customMiddleware.TokenValidate).Put("/{id}", blogHandler.Update)
			r.With(customMiddleware.TokenValidate).Delete("/{id}", blogHandler.Delete)
		})
		r.Route("/comments", func(r chi.Router) {
			r.Use(customMiddleware.TokenValidate)
			r.Post("/", commentHandler.Create)
			r.Put("/{id}", commentHandler.Update)
			r.Delete("/{id}", commentHandler.Delete)
		})
	})

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Printf(err.Error())
	}
}
