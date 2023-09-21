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
	db             = config.OpenDatabase()
	userRepository = repository.NewUserRepository(db)
	userService    = service.NewUserService(userRepository)
	authService    = service.NewAuthService(userRepository)
	authHandler    = handler.NewAuthHandler(authService)
	userHandler    = handler.NewUserHandler(userService)
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
		r.Post("/login", authHandler.Login)
		r.Post("/register", authHandler.Register)

		r.Route("/account", func(r chi.Router) {
			r.Use(customMiddleware.TokenValidate)
			r.Get("/", userHandler.Profile)
			r.Post("/update", userHandler.Update)
			r.Post("/update-password", userHandler.UpdatePassword)
		})
	})

	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Printf(err.Error())
	}
}
