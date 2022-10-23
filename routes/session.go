package routes

import (
	"fmt"

	"github.com/Bendimester23/go-auth-server/controllers"
	"github.com/Bendimester23/go-auth-server/models"
	"github.com/gofiber/fiber/v2"
)

var sessionController = &controllers.SessionController{}

func HandleSession_Join(c *fiber.Ctx) error {
	req := new(models.JoinRequest)
	c.BodyParser(req)
	if err := req.Validate(); len(err) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"errors": err,
		})
	}

	err := sessionController.Join(req)
	if err != nil {
		return c.Status(err.Code).SendString(err.Message)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func HandleSession_HasJoined(c *fiber.Ctx) error {
	fmt.Println(c.Query("serverId"))
	res, err := sessionController.HasJoined(c.Query("username"), c.Query("serverId"))
	if err != nil {
		return c.Status(err.Code).SendString(err.Message)
	}

	return c.JSON(res)
}

func HandleSession_UuidToProfile(c *fiber.Ctx) error {
	res, err := sessionController.PlayerProfile(c.Params("uuid"))
	if err != nil {
		return c.Status(err.Code).SendString(err.Message)
	}

	return c.JSON(res)
}
