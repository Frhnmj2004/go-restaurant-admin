package main

import (
	"log"
	"os"

	"github.com/Frhnmj2004/restaurant-admin/database"
	"github.com/Frhnmj2004/restaurant-admin/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := database.NewConnection(config)
	if err != nil {
		log.Fatal("could not connect to the database: ", err)
	}

	err = database.MigrateDB(db)
	if err != nil {
		log.Fatal("failed to migrate the database: ", err)
	}

	r := routes.Repository{DB: db}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")

}
