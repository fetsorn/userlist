package responses

import (
	"userlist/models"
)

type ListallResponse struct {
	Httpstatus bool          `json:"success"`
	Data       *models.Users `json:"data"`
}

func Listalluser(users *models.Users) *ListallResponse {
	resp := &ListallResponse{Data: users}
	return resp
}
