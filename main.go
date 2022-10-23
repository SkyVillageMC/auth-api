package main

import (
	"log"
	"net/http"

	"github.com/Bendimester23/go-auth-server/db"
	"github.com/Bendimester23/go-auth-server/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func loadConfig() {
	viper.AddConfigPath("/etc/skyvillage/auth")
	viper.AddConfigPath(".")

	viper.SetDefault("Listen", ":8080")
	viper.SetDefault("StaticUrl", "http://localhost:8080/textures")
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	log.Println("Starting...")
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	db.Connect()
	defer db.Disconnect()

	loadConfig()
	setupRoutes(app)

	if err := app.Listen(viper.GetString("Listen")); err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/authenticate", routes.HandleAuth_Authenticate)
	auth.Post("/refresh", routes.HandleAuth_Refresh)
	auth.Post("/validate", routes.HandleAuth_Validate)
	auth.Post("/signout", routes.HandleAuth_SignOut)
	auth.Post("/invalidate", routes.HandleAuth_Invalidate)

	session := app.Group("/session/minecraft")
	session.Post("/join", routes.HandleSession_Join)
	session.Get("/hasJoined", routes.HandleSession_HasJoined)
	session.Get("/profile/:uuid", routes.HandleSession_UuidToProfile)

	api := app.Group("/api")
	api.Get("/player/attributes", routes.HandleApi_Attributes)

	app.Use("/textures", filesystem.New(filesystem.Config{
		Root:         http.Dir("./static"),
		NotFoundFile: "skins/steve.png",
		MaxAge:       3600,
		Browse:       false,
	}))
}
