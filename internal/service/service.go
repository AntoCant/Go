package service

import (
	"Api-users/internal/model"
	"Api-users/internal/repository"
)

func GetAllUsers() ([]model.User, error) {
	return repository.GetAllUsers()
}

func CreateUser(user model.User) error {
	return repository.CreateUser(user)
}
