package service

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func Run(config *AppConfig) {
	router := fiber.New()
	if config.Middlewares != nil {
		for _, middleware := range * config.Middlewares {
			router.Use(middleware)
		}
	}

	if config.Controllers != nil {
		for _, controller := range *config.Controllers {
			controller.Register(router)
		}
	}


	if err := router.Listen(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)); err != nil {
		panic(err)
	}
}