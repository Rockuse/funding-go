package user

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	// Get(param int) ([]User, error)
	// GetAll(user []User) ([]User, error)
	// Delete(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(ID int) (User, error)
	Update(user User) (User, error)
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

func (r *repository) FindByEmail(email string) (User, error) {
	var userData User
	err := r.db.Find(&userData, "email=?", email).Error
	if err != nil {
		return userData, err
	}
	return userData, nil
}

func (r *repository) FindById(ID int) (User, error) {
	var user User
	err := r.db.Find(&user, "id=?", ID).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Model(&user).Updates(User{Avatar_file_name: user.Avatar_file_name, ModifiedDate: time.Now(), ModifiedBy: strconv.Itoa(user.Id)}).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// func (r *repository) GetAll(users []User) ([]User, error) {
// 	err := r.db.Find(&users).Error
// 	if err != nil {
// 		return users, err
// 	}
// 	return users, nil
// }

// func (r *repository) Delete(users User, param int) (User, error) {
// 	err := r.db.Delete(&users, "id=?", param)
// 	if err != nil {
// 		return users, nil
// 	}
// 	return users, nil
// }
