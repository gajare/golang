package service

import "Amazon/user/internal/models"

// The service layer typically encapsulates business logic that is reusable across different controllers or even across multiple microservices.

func CreateUser(user models.User) (*models.User, error) {
	return &user, nil
}

func UpadateUser(user_id string, user models.User) (*models.User,error) {
	return &user, nil
}

func DeleteUser(user models.User)(*models.User,error){
	return &user,nil
}

func ListUsers(user models.User)(*models.User,error){
	return &user,nil
}
