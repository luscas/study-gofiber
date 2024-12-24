package main

import (
	"fmt"
	"log"
	"os"

	"api.droppy.com.br/internal/middlewares"
	"api.droppy.com.br/internal/routes"
	"api.droppy.com.br/pkg/database"
	"api.droppy.com.br/pkg/jwt"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env:", err)
	}

	jwt.InitSecretKey()

	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatal("Erro ao carregar o db:", err)
	}

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(middlewares.Recover())
	app.Use(middlewares.Logger())

	app.Static("/", "./public")

	routes.SetupRoutes(
		app,
		db,
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	if err := app.Listen(":" + port); err != nil {
		fmt.Println(err)
	}
}
