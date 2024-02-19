package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kauri646/be-idealibs/internal/auth"
	"github.com/kauri646/be-idealibs/internal/users/handlers"
	"github.com/kauri646/be-idealibs/middleware"
)

func RouteInit(r *fiber.App) {
	r.Get("/auth/google", auth.GoogleOauth)
	r.Get("/auth/google/callback", auth.GoogleCallback)
	r.Get("/profile", auth.Profil)

	r.Post("/login", handlers.LoginHandler)
	r.Post("/register", handlers.RegisterHandler)

	r.Get("/user", middleware.Auth, handlers.UserHandlerGetAll)
	r.Get("/user/:id", handlers.UserHandlerGetById)
	r.Post("/user", handlers.UserHandlerCreate)
	r.Put("/user/:id", handlers.UserHandlerUpdate)
	r.Put("/user/:id/update-email", handlers.UserHandlerUpdateEmail)
	r.Delete("/user/:id", handlers.UserHandlerDelete)
}
