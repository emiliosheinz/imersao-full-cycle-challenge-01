package usecase

import "github.com/emiliosheinz/imersao-full-cycle-challenge-01/domain/model"

type ProductUseCase struct {
	ProductRepository model.ProductRepositoryInterface
}

func (useCase ProductUseCase) CreateProduct(name string, description string, price float32) (*model.Product, error) {
	product, err := model.NewProduct(name, description, price)

	if err != nil {
		return nil, err
	}

	return useCase.ProductRepository.Create(product)
}

func (useCase ProductUseCase) FindProducts() (*[]model.Product, error) {
	return useCase.ProductRepository.FindAll()
}
