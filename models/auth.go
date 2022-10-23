package models

import "github.com/go-playground/validator/v10"

var Validator = validator.New()

type AuthenticateRequest struct {
	Username    string `json:"username" validate:"required,min=2,max=40"`
	Password    string `json:"password" validate:"required,min=3,max=30"`
	ClientToken string `json:"clientToken" validate:"required,min=30,max=40"`
	RequestUser bool   `json:"requestUser"`
}

type IError struct {
	Field string
	Tag   string
	Value string
}

func (r *AuthenticateRequest) Validate() []IError {
	var errors []IError

	err := Validator.Struct(r)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, el)
		}
	}
	return errors
}

func (r *RefreshRequest) Validate() []IError {
	var errors []IError

	err := Validator.Struct(r)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, el)
		}
	}
	return errors
}

func (r *SignOutRequest) Validate() []IError {
	var errors []IError

	err := Validator.Struct(r)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, el)
		}
	}
	return errors
}

type AuthenticateResponse struct {
	User              User      `json:"user"`
	ClientToken       string    `json:"clientToken"`
	AccessToken       string    `json:"accessToken"`
	AvailableProfiles []Profile `json:"availableProfiles"`
	SelectedProfile   Profile   `json:"selectedProfile"`
}

type Profile struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type User struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

type RefreshRequest struct {
	AccessToken string `json:"accessToken" validate:"requird,min=16,max=32"`
	ClientToken string `json:"clientToken" validate:"required,min=30,max=40"`

	RequestUser bool `json:"requestUser"`
}

type SignOutRequest struct {
	Username string `json:"username" validate:"required,min=2,max=40"`
	Password string `json:"password" validate:"required,min=3,max=30"`
}
