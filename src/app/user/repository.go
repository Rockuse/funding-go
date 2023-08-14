package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	Get(user User, param int) ([]User, error)
	GetAll() ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// func (r *repository) Get(user User, param int) ([]User, error) {
// 	var users []User
// 	users, err := r.db.Find(&user, "id=?", string(param))
// 	if err != nil {
// 		return users, err
// 	}
// 	return users, nil
// }
