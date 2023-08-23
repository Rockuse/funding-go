package user

import (
	"time"
)

type User struct {
	Id               int
	Name             string
	Occupation       string
	Email            string
	Password         string
	Avatar_file_name string
	Role             string
	Token            string
	CreatedDate      time.Time
	ModifiedDate     time.Time
	CreatedBy        string
	ModifiedBy       string
}

type RegisterInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
