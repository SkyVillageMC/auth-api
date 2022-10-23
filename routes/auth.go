package routes

import (
	"github.com/Bendimester23/go-auth-server/controllers"
	"github.com/Bendimester23/go-auth-server/models"
	"github.com/gofiber/fiber/v2"
)

var authController = &controllers.AuthController{}

func HandleAuth_Authenticate(c *fiber.Ctx) error {
	req := new(models.AuthenticateRequest)
	c.BodyParser(req)
	if err := req.Validate(); len(err) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"errors": err,
		})
	}

	res, err := authController.Authenticate(req)
	if err != nil {
		return c.Status(err.Code).SendString(err.Message)
	}

	return c.JSON(res)
}

func HandleAuth_Refresh(c *fiber.Ctx) error {
	req := new(models.RefreshRequest)
	c.BodyParser(req)
	if err := req.Validate(); len(err) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"errors": err,
		})
	}

	res, err := authController.Refresh(req)
	if err != nil {
		return c.Status(err.Code).SendString(err.Message)
	}

	return c.JSON(res)
}

func HandleAuth_Validate(c *fiber.Ctx) error {
	req := new(models.RefreshRequest)
	c.BodyParser(req)
	if err := req.Validate(); len(err) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"errors": err,
		})
	}

	err := authController.Validate(req)
	if err != nil {
		return c.Status(err.Code).SendString(err.Message)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func HandleAuth_SignOut(c *fiber.Ctx) error {
	req := new(models.SignOutRequest)
	c.BodyParser(req)
	if err := req.Validate(); len(err) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"errors": err,
		})
	}

	err := authController.SignOut(req)
	if err != nil {
		return c.Status(err.Code).SendString(err.Message)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func HandleAuth_Invalidate(c *fiber.Ctx) error {
	req := new(models.RefreshRequest)
	c.BodyParser(req)
	if err := req.Validate(); len(err) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"errors": err,
		})
	}

	err := authController.Invalidate(req)
	if err != nil {
		return c.Status(err.Code).SendString(err.Message)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
