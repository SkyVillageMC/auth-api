package models

import "github.com/go-playground/validator/v10"

type JoinRequest struct {
	AccessToken     string `json:"accessToken" validate:"required,min=30,max=40"`
	SelectedProfile string `json:"selectedProfile"`
	ServerId        string `json:"serverId"`
}

func (r *JoinRequest) Validate() []IError {
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
