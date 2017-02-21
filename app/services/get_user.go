package services

import (
	"errors"
	"log"

	"generator/app/domain"
	"generator/app/storage/repository/generator"
)

//GetService service to get user
type GetUserService struct{}

//NewGetService construct service to get user
func NewGetUserService() GetUserService {
	return GetUserService{}
}

func (this GetUserService) GetUser() (domain.User, error) {
	user, err := generator.NewUser()
	if err != nil {
		return domain.User{}, err
	}

	if !user.Verify() {
		log.Printf("Incorrect user generated %#+v", user)
		return domain.User{}, errors.New("Internal error")
	}

	return *user, nil
}
