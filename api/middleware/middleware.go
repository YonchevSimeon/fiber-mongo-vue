package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(
		cors.New(),
		logger.New(),
	)
}

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey:   []byte("secret"),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	return jwtware.New(config)
}

func jwtError(ctx *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return ctx.Status(400).JSON(fiber.Map{
			"ok":  false,
			"err": err.Error(),
		})
	}

	return ctx.Status(404).JSON(fiber.Map{
		"ok":  false,
		"err": err.Error(),
	})
}
