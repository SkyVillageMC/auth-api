package controllers

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/Bendimester23/go-auth-server/db"
	"github.com/Bendimester23/go-auth-server/models"
	"github.com/lucsky/cuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
}

var ctx = context.TODO()

type ErrorResponse struct {
	Code    int
	Message string
}

func (a *AuthController) Authenticate(data *models.AuthenticateRequest) (*models.AuthenticateResponse, *ErrorResponse) {
	res, err := db.DB.User.FindFirst(
		db.User.Email.Equals(data.Username),
	).Exec(ctx)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, &ErrorResponse{
				Code:    404,
				Message: "Email or Password invalid",
			}
		}
		return nil, &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(data.Password)) != nil {
		return nil, &ErrorResponse{
			Code:    404,
			Message: "Email or Password invalid",
		}
	}

	accessToken := a.makeAccessToken(res)

	_, err = db.DB.User.FindMany(
		db.User.ID.Equals(res.ID),
	).Update(
		db.User.ClientToken.Set(data.ClientToken),
		db.User.AccessToken.Set(accessToken),
	).Exec(ctx)
	if err != nil {
		return nil, &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}

	profile := models.Profile{
		Name: res.Username,
		ID:   res.UUID,
	}
	if data.RequestUser {
		return &models.AuthenticateResponse{
			User: models.User{
				Username: res.Email,
				ID:       res.ID,
			},
			ClientToken:       data.ClientToken,
			AccessToken:       accessToken,
			AvailableProfiles: []models.Profile{profile},
			SelectedProfile:   profile,
		}, nil
	}

	return &models.AuthenticateResponse{
		ClientToken:       data.ClientToken,
		AccessToken:       accessToken,
		AvailableProfiles: []models.Profile{profile},
		SelectedProfile:   profile,
	}, nil
}

func (a *AuthController) Refresh(data *models.RefreshRequest) (*models.AuthenticateResponse, *ErrorResponse) {
	res, err := db.DB.User.FindFirst(
		db.User.ClientToken.Equals(data.ClientToken),
		db.User.AccessToken.Equals(data.AccessToken),
	).Exec(ctx)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return nil, &ErrorResponse{
				Code:    404,
				Message: "Email or Password invalid",
			}
		}
		return nil, &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}

	accessToken := a.makeAccessToken(res)
	_, err = db.DB.User.FindMany(
		db.User.ID.Equals(res.ID),
	).Update(
		db.User.AccessToken.Set(accessToken),
	).Exec(ctx)

	if err != nil {
		return nil, &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}

	if data.RequestUser {
		return &models.AuthenticateResponse{
			ClientToken: data.ClientToken,
			AccessToken: data.AccessToken,
			SelectedProfile: models.Profile{
				Name: res.Username,
				ID:   res.UUID,
			},
			User: models.User{
				Username: res.Email,
				ID:       res.ID,
			},
		}, nil
	}

	return &models.AuthenticateResponse{
		ClientToken: data.ClientToken,
		AccessToken: data.AccessToken,
		SelectedProfile: models.Profile{
			Name: res.Username,
			ID:   res.UUID,
		},
	}, nil
}

func (a *AuthController) Validate(data *models.RefreshRequest) *ErrorResponse {
	_, err := db.DB.User.FindFirst(
		db.User.ClientToken.Equals(data.ClientToken),
		db.User.AccessToken.Equals(data.AccessToken),
	).Exec(ctx)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return &ErrorResponse{
				Code:    404,
				Message: "Email or Password invalid",
			}
		}
		return &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}
	return nil
}

func (a *AuthController) SignOut(data *models.SignOutRequest) *ErrorResponse {
	res, err := db.DB.User.FindFirst(
		db.User.Email.Equals(data.Username),
	).Exec(ctx)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return &ErrorResponse{
				Code:    404,
				Message: "Email or Password invalid",
			}
		}
		return &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}
	if bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(data.Password)) != nil {
		return &ErrorResponse{
			Code:    404,
			Message: "Email or Password invalid",
		}
	}
	_, err = db.DB.User.FindMany(
		db.User.ID.Equals(res.ID),
	).Update(
		db.User.AccessToken.Set(""),
		db.User.ClientToken.Set(""),
	).Exec(ctx)

	if err != nil {
		return &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}
	return nil
}

func (a *AuthController) Invalidate(data *models.RefreshRequest) *ErrorResponse {
	res, err := db.DB.User.FindFirst(
		db.User.ClientToken.Equals(data.ClientToken),
		db.User.AccessToken.Equals(data.AccessToken),
	).Exec(ctx)
	if err != nil {
		if errors.Is(err, db.ErrNotFound) {
			return &ErrorResponse{
				Code:    404,
				Message: "Email or Password invalid",
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
		db.User.AccessToken.Set(""),
		db.User.ClientToken.Set(""),
	).Exec(ctx)

	if err != nil {
		return &ErrorResponse{
			Code:    500,
			Message: "Server error",
		}
	}
	return nil
}

func (a *AuthController) makeAccessToken(user *db.UserModel) string {
	hash := md5.New()
	hash.Write([]byte(user.Username))
	hash.Write([]byte(cuid.New()))
	return hex.EncodeToString(hash.Sum([]byte{}))
}
