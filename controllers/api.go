package controllers

import (
	"errors"

	"github.com/Bendimester23/go-auth-server/db"
)

type ApiController struct {
}

type PlayerAttributes struct {
	AllowRealms      bool
	AllowMultiplayer bool
	AllowChat        bool
}

func (a *ApiController) Attributes(accessToken string) (*PlayerAttributes, *ErrorResponse) {
	res, err := db.DB.User.FindFirst(
		db.User.AccessToken.Equals(accessToken),
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
	return &PlayerAttributes{
		AllowRealms:      res.AllowRealms,
		AllowMultiplayer: res.AllowMultiplayer,
		AllowChat:        res.AllowChat,
	}, nil
}
