package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/smochii/go-clean-architecture/handler"
)

type Route struct {
	middlewareHandler handler.MiddlewareHandler
	authHandler       handler.AuthHandler
	userHandler       handler.UserHandler
}

func NewRoute(middlewareHandler handler.MiddlewareHandler, authHandler handler.AuthHandler, userHandler handler.UserHandler) *Route {
	return &Route{
		middlewareHandler: middlewareHandler,
		authHandler:       authHandler,
		userHandler:       userHandler,
	}
}

func (r *Route) Register(app *fiber.App) {
	app.Get("", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Post("/auth/token", r.authHandler.GenerateToken)
	app.Post("/users", r.userHandler.Create)

	auth := app.Group("")
	auth.Use(r.middlewareHandler.AuthenticateMiddleware())
	auth.Get("/me", r.userHandler.Me)
}
