package view

import (
	"github.com/gustavoeyros/golang-project/src/controllers/models/response"
	"github.com/gustavoeyros/golang-project/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
