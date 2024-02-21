package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kauri646/be-idealibs/internal/auth"
	handlersimages "github.com/kauri646/be-idealibs/internal/images/handlers"
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

	r.Get("/images", handlersimages.ImageHandlerGetAll)
	r.Get("/images/:id", handlers.UserHandlerGetById)
	r.Post("/images", handlersimages.ImageHandlerCreate)
	//r.Put("/images/:id", handlers.UserHandlerUpdate)
	//r.Put("/images/:id/update-email", handlers.UserHandlerUpdateEmail)
	r.Delete("/images/:id", handlers.UserHandlerDelete)
}
