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
