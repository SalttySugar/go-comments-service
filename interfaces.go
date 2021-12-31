package service

import (
	"github.com/gofiber/fiber/v2"
)

type Controller interface {
	Register(a *fiber.App)
}

type CommentsService interface {
	FindById(string) Comment
	FindAll() []Comment
	DeleteById(string)
	Create(CommentDto) Comment
	Update(string, string) Comment
}
