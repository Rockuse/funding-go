package user

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input EmailInput) (bool, error)
	UpdateAvatar(ID int, filepath string) (User, error)
	GetUserById(ID int) (User, error)
	UpdateToken(ID int, token string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)
	user.Role = "user"
	user.CreatedDate = time.Now()
	user.ModifiedDate = time.Now()
	user.CreatedBy = "Fahmi"
	user.ModifiedBy = "Fahmi"
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password
	userData, err := s.repository.FindByEmail(email)
	if err != nil {
		return userData, err
	}
	if userData.Id == 0 {
		return userData, errors.New("User Not Found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		return User{}, err
	}
	return userData, nil
}

func (s *service) IsEmailAvailable(input EmailInput) (bool, error) {
	email := input.Email
	userData, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if userData.Email != "" {
		return false, nil
	}
	return true, nil
}

func (s *service) UpdateAvatar(ID int, filepath string) (User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("User Not Found")
	}
	user.Avatar_file_name = filepath
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

func (s *service) GetUserById(ID int) (User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) UpdateToken(ID int, token string) (User, error) {
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}
	user.Token = token
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}
