package infrastructure

import (
	x_app "go-fiber-template/internal/app"
	"go-fiber-template/internal/auth"
	"go-fiber-template/internal/docs"
	"go-fiber-template/internal/schedule"
	"go-fiber-template/internal/user"
	"go-fiber-template/lib/common"

	"github.com/gofiber/fiber/v2"
)

func registerRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	x_app.NewHttpHandler(api)
	docs.NewHttpHandler(api.Group("/docs"))
	auth.NewHttpHandler(api.Group("/auth"), authService)
	user.NewHttpHandler(api.Group("/users"), userService)
	schedule.NewHttpHandler(api.Group("/schedules"), scheduleService)
	app.Use(common.NotFoundHandler)
}
