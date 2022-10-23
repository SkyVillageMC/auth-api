package controllers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"time"

	"github.com/Bendimester23/go-auth-server/db"
	"github.com/Bendimester23/go-auth-server/models"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type SessionController struct {
}

func (s *SessionController) Join(data *models.JoinRequest) *ErrorResponse {
	res, err := db.DB.User.FindFirst(
		db.User.UUID.Equals(data.SelectedProfile),
		db.User.AccessToken.Equals(data.AccessToken),
	).Exec(ctx)

	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return &ErrorResponse{
				Code:    403,
				Message: "Invalid session",
			}
		}
		return &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}

	_, err = db.DB.User.FindMany(
		db.User.ID.Equals(res.ID),
	).Update(
		db.User.CurrentServer.Set(data.ServerId),
	).Exec(ctx)

	if err != nil {
		return &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}
	return nil
}

func (s *SessionController) HasJoined(username string, serverId string) (fiber.Map, *ErrorResponse) {
	res, err := db.DB.User.FindFirst(
		db.User.Username.Equals(username),
		db.User.CurrentServer.Equals(serverId),
	).Exec(ctx)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, &ErrorResponse{
				Code:    403,
				Message: "Invalid session",
			}
		}
		return nil, &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}

	return s.CreateProfile(res), nil
}

func (s *SessionController) CreateProfile(res *db.UserModel) fiber.Map {
	skin := "steve.png"
	if r, ok := res.SkinID(); ok {
		skin = r + ".png"
	}

	blob := fiber.Map{
		"timestamp":   time.Now().UnixMilli(),
		"profileId":   res.UUID,
		"profileName": res.Username,
		"textures": fiber.Map{
			"SKIN": fiber.Map{
				"url": viper.GetString("StaticUrl") + "/skins/" + skin,
			},
		},
	}

	if r, ok := res.CapeID(); ok {
		blob["textures"].(fiber.Map)["CAPE"] = fiber.Map{
			"url": viper.GetString("StaticUrl") + "/capes/" + r + ".png",
		}
	}

	raw, _ := json.Marshal(blob)

	textures := base64.StdEncoding.EncodeToString(raw)

	return fiber.Map{
		"id":   res.UUID,
		"name": res.Username,
		"properties": []fiber.Map{
			{
				"name":  "textures",
				"value": textures,
			},
		},
	}
}

func (s *SessionController) PlayerProfile(uuid string) (fiber.Map, *ErrorResponse) {
	res, err := db.DB.User.FindFirst(
		db.User.UUID.Equals(uuid),
	).Exec(ctx)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, &ErrorResponse{
				Code:    404,
				Message: "Player not found",
			}
		}
		return nil, &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}

	return s.CreateProfile(res), nil
}
