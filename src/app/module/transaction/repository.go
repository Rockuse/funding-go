package transaction

import "gorm.io/gorm"

type Repository interface {
	Add(transaction Transaction) (Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Add() {

}
