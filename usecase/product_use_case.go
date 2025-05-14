package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

// trata as regras de negocio da rota getProducts

func (p *ProductUseCase) GetProducts() ([]model.Product, error) {
	return p.repository.GetProducts()
}
