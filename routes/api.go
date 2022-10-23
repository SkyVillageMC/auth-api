package routes

import (
	"github.com/Bendimester23/go-auth-server/controllers"
	"github.com/gofiber/fiber/v2"
)

var apiController = &controllers.ApiController{}

func HandleApi_Attributes(c *fiber.Ctx) error {
	accessToken := c.Get("Authorization")[7:]

	res, err := apiController.Attributes(accessToken)
	if err != nil {
		return c.Status(err.Code).SendString(err.Message)
	}

	return c.JSON(fiber.Map{
		"privileges": fiber.Map{
			"onlineChat": fiber.Map{
				"enabled": res.AllowChat,
			},
			"multiplayerServer": fiber.Map{
				"enabled": res.AllowMultiplayer,
			},
			"multiplayerRealms": fiber.Map{
				"enabled": res.AllowRealms,
			},
			"telemetry": fiber.Map{
				"enabled": false,
			},
		},
		"profanityFilterPreferences": fiber.Map{
			"profanityFilterOn": false,
		},
		"banStatus": fiber.Map{
			"bannedScopes": fiber.Map{},
		},
	})
}
