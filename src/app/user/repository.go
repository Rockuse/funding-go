package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	Get(user User) ([]User, error)
	GetAll(user User) ([]User, error)
	Delete(user User) (User, error)
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

func (r *repository) Get(param int) ([]User, error) {
	var users []User
	err := r.db.Find(&users, "id = ?", param).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (r *repository) GetAll() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
