package useCase

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

func (p *ProductUseCase) CreateProduct(product model.Product) (model.Product, error) {
	productID, err := p.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = productID
	return product, nil
}

func (p *ProductUseCase) GetProductsById(id_product int) (model.Product, error) {
	product, err := p.repository.FindProductsById(id_product)

	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}
