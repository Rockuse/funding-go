package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
<<<<<<< HEAD
	Get(user User) ([]User, error)
	GetAll(user User) ([]User, error)
	Delete(user User) (User, error)
=======
	// Get(param int) ([]User, error)
	GetAll(user []User) ([]User, error)
	// Delete(user User) (User, error)
>>>>>>> 0421883400e7d8f9ec598f63f7e1966b6af40f53
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
<<<<<<< HEAD
	err := r.db.Find(&users, "id = ?", param).Error
=======
	err := r.db.Find(&users, "id=?", param).Error
>>>>>>> 0421883400e7d8f9ec598f63f7e1966b6af40f53
	if err != nil {
		return users, err
	}
	return users, nil
}

<<<<<<< HEAD
func (r *repository) GetAll() ([]User, error) {
	var users []User
=======
func (r *repository) GetAll(users []User) ([]User, error) {
>>>>>>> 0421883400e7d8f9ec598f63f7e1966b6af40f53
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
<<<<<<< HEAD
=======

func (r *repository) Delete(users User, param int) (User, error) {
	err := r.db.Delete(&users, "id=?", param)
	if err != nil {
		return users, nil
	}
	return users, nil
}
>>>>>>> 0421883400e7d8f9ec598f63f7e1966b6af40f53
