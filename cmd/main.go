package main

import (
	service "blog.com/services/comments/v1"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config := service.LoadConfig()
	service.Run(&service.AppConfig{
		Server: config.Server,
		Middlewares: &[]fiber.Handler{
			cache.New(config.Caching),
			recover2.New(),
			compress.New(),
			logger.New(),
		},
		Controllers: &[]service.Controller{
			service.NewCommentsController(service.NewGormCommentsService(service.Connect(config.Database))),
		},
	})
}
