package user

import "time"

type User struct {
	Id           int
	Name         string
	Occupation   string
	Email        string
	Password     string
	Avatar       string
	Role         string
	Token        string
	CreatedDate  time.Time
	ModifiedDate time.Time
	CreatedBy    string
	ModifiedBy   string
}
