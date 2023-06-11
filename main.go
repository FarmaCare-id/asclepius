package main

import (
	"farmacare/di"
	"farmacare/docs"
	"farmacare/infrastructure"
	"log"

	"farmacare/shared/config"

	"github.com/gofiber/fiber/v2"
)

// @title FarmaCare Server
// @version 1.0
// @description API definition for FarmaCare Server
// @host localhost:8000
// @BasePath /
func main() {
	container := di.Container

	err := container.Invoke(func(http *fiber.App, env *config.EnvConfig, holder infrastructure.Holder) error {
		infrastructure.Routes(http, holder)

		docs.SwaggerInfo.Host = env.HOST

		err := http.Listen(":" + env.PORT)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalf("error when starting http server: %s", err.Error())
	}
}