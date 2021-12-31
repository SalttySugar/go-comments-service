package service

import "github.com/gofiber/fiber/v2"

type CommentsController struct {
	service CommentsService
}


func (controller *CommentsController) Register(router *fiber.App) {
	router.Get("/api/v1/comments/", controller.findAllComments())
	router.Get("/api/v1/comments/:id", controller.getCommentById())
	router.Post("/api/v1/comments/", controller.createComment())
	router.Put("/api/v1/comments/:id", controller.updateComment())
	router.Delete("/api/v1/comments/:id", controller.deleteCommentById())
}




func (controller *CommentsController) findAllComments() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON(controller.service.FindAll())
	}
}

func (controller *CommentsController) getCommentById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id");
		return c.JSON(controller.service.FindById(id))
	}
}

func (controller *CommentsController) createComment() fiber.Handler {

	return func(c *fiber.Ctx) error {
		payload := new(CommentDto)
		if err := c.BodyParser(payload); err != nil {
			return err
		}
		return c.JSON(controller.service.Create(*payload))
	}
}

func (controller *CommentsController) updateComment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		payload := new(CommentDto)
		if err := c.BodyParser(payload); err != nil {
			return err
		}
		return c.JSON(controller.service.Update(id, payload.Message))
	}
}

func (controller *CommentsController) deleteCommentById() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		controller.service.DeleteById(id)
		return c.Send(nil)
	}
}

var controller *CommentsController

func NewCommentsController(service CommentsService) *CommentsController {
	if controller != nil {return controller}
	controller = &CommentsController{service: service}
	return controller;
}