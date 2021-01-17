package responses

import (
	"userlist/models"
)

type ListOneResponse struct {
	Httpstatus bool         `json:"success"`
	Data       *models.User `json:"data"`
}

type ListAllResponse struct {
	Httpstatus bool          `json:"success"`
	Data       *models.Users `json:"data"`
}

func ListOne(user *models.User) *ListOneResponse {
	r := &ListOneResponse{Data: user}
	return r
}

func ListAll(users *models.Users) *ListAllResponse {
	r := &ListAllResponse{Data: users}
	return r
}
