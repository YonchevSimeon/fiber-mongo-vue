package routes

import (
	"github.com/YonchevSimeon/fiber-mongodb-vue/controllers"
	"github.com/YonchevSimeon/fiber-mongodb-vue/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetUpPublicRoutes(app *fiber.App) {
	route := app.Group("/api")

	route.Post("/login", controllers.Login)
	route.Post("/register", controllers.Register)
}

func SetUpPrivateRoutes(app *fiber.App) {
	route := app.Group("/api")

	route.Get("/posts", middleware.JWTProtected(), controllers.GetAll)
	route.Get("/posts/:id", middleware.JWTProtected(), controllers.GetById)
	route.Post("/posts", middleware.JWTProtected(), controllers.Create)
	route.Patch("/posts/:id", middleware.JWTProtected(), controllers.Update)
	route.Delete("/posts/:id", middleware.JWTProtected(), controllers.Delete)

	route.Get("/user/:id/posts", middleware.JWTProtected(), controllers.GetByUserId)
}

func NotFoundRoute(app *fiber.App) {
	app.Use(
		func(ctx *fiber.Ctx) error {
			return ctx.Status(404).JSON(fiber.Map{
				"ok":    false,
				"error": "Sorry, the page you are looking for was NOT found!",
			})
		},
	)
}
