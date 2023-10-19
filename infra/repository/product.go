package repository

import (
	"github.com/emiliosheinz/imersao-full-cycle-challenge-01/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionProductRepositoryDb struct {
	Db *gorm.DB
}

func (repo TransactionProductRepositoryDb) FindAll() (*[]model.Product, error) {
	var products []model.Product

	repo.Db.Find(&products)

	return &products, nil
}

func (repo TransactionProductRepositoryDb) Create(product *model.Product) (*model.Product, error) {
	err := repo.Db.Create(&product).Error

	if err != nil {
		return nil, err
	}

	return product, nil
}
