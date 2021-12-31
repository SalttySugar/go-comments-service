package service

import "github.com/gofiber/fiber/v2"

func ErrCommentNotFound(id string) error{
	return fiber.NewError(fiber.StatusNotFound, "{\"error\": \"not found \"}");
}