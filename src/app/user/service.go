package user

import "golang.org/x/crypto/bcrypt"

type RegisterInput struct {
	Name       string
	Occupation string
	Email      string
	Password   string
}

type Service interface {
	RegisterUser(input RegisterInput) (User, error)
	// GetAllUser(users User) ([]User, error)
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
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) GetAllUser(users []User) ([]User, error) {
	users, err := s.repository.GetAll(users)
	if err != nil {
		return users, nil
	}
	return users, nil
}
