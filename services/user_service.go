package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (service *Service) AddUserService(u *models.User) (res *mongo.InsertOneResult, err error) {
	result, err := service.userdao.AddUser(u)
	if err != nil {
		return result, err
	}
	return result, err
}

func (service *Service) DeleteUserService(id primitive.ObjectID) (res *mongo.DeleteResult, err error) {
	filter := bson.M{"_id": id}
	res, err = service.userdao.DeleteUser(filter)
	if err != nil {
		return res, err
	}
	return res, err
}

func (service *Service) GetUserService(id primitive.ObjectID) (user models.User, err error) {
	filter := bson.M{"_id": id}
	result := service.userdao.GetUser(filter)
	if result.Err() != nil {
		return user, result.Err()
	}
	err = result.Decode(&user)
	return user, err
}

func (service *Service) GetUserByFirstService(p string) (user models.User, err error) {
	filter := bson.M{"first": p}
	result := service.userdao.GetUser(filter)
	if result.Err() != nil {
		return user, result.Err()
	}
	err = result.Decode(&user)
	return user, err
}

func (service *Service) GetUserByLastService(p string) (user models.User, err error) {
	filter := bson.M{"last": p}
	result := service.userdao.GetUser(filter)
	if result.Err() != nil {
		return user, result.Err()
	}
	err = result.Decode(&user)
	return user, err
}

func (service *Service) GetUserByCityService(p string) (user models.User, err error) {
	filter := bson.M{"city": p}
	result := service.userdao.GetUser(filter)
	if result.Err() != nil {
		return user, result.Err()
	}
	err = result.Decode(&user)
	return user, err
}

func (service *Service) GetUserByCountryService(p string) (user models.User, err error) {
	filter := bson.M{"country": p}
	result := service.userdao.GetUser(filter)
	if result.Err() != nil {
		return user, result.Err()
	}
	err = result.Decode(&user)
	return user, err
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
