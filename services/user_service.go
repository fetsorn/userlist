package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"userlist/dao"
	"userlist/models"
)

type Service struct {
	userdao dao.UserDaoInterface
}

var service *Service

func NewService() *Service {
	service = &Service{dao.NewDao()}
	return service
}

func (service *Service) GetUsersService() (Users []models.User, err error) {
	cur, errs := service.userdao.GetUsers(bson.M{})
	if errs != nil {
		return Users, errs
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var user models.User

		err := cur.Decode(&user)
		if err != nil {
			return Users, err
		}
		Users = append(Users, user)
	}
	return Users, err
}
