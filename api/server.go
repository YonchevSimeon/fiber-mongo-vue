package main

import (
	"log"
	"os"

	"github.com/Kamva/mgm/v2"
	"github.com/YonchevSimeon/fiber-mongodb-vue/controllers"
	"github.com/gofiber/fiber"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := fiber.New()

	// Routes
	app.Get("/api/posts", controllers.GetAll)
	app.Get("/api/posts/:id", controllers.GetById)
	app.Post("/api/posts", controllers.Create)
	app.Delete("/api/posts/:id", controllers.Delete)

	err := app.Listen(8000)

	if err != nil {
		panic(err)
	}
}

func init() {
	connectionString := getDbConnectionString("MONGODB_CONNECTION_STRING")

	if len(connectionString) == 0 {
		connectionString = "mongodb://localhost:27017"
	}

	err := mgm.SetDefaultConfig(nil, "forum", options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}
}

func getDbConnectionString(key string) string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
