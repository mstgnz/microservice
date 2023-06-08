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
	db             = config.OpenDatabase()
	userRepository = repository.UserRepository(db)
	userService    = service.UserService(userRepository)
	authService    = service.AuthService(userRepository)
	authHandler    = handler.AuthHandler(authService)
	userHandler    = handler.UserHandler(userService)
)

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
		r.Post("/login", authHandler.Login)
		r.Post("/register", authHandler.Register)

		r.Route("/user", func(r chi.Router) {
			r.Use(customMiddleware.TokenValidate)
			r.Get("/profile", userHandler.Profile)
			r.Post("/update", userHandler.Update)
		})
	})

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Printf(err.Error())
	}
}
