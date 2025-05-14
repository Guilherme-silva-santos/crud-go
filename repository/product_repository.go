package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	// query qye será feita no bano

	query := "SELECT id, product_name, price FROM product"

	// executando a query
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err := rows.Scan(&productObj.ID, &productObj.Name, &productObj.Price)

		if err != nil {
			fmt.Println("Error executing query:", err)
			return []model.Product{}, err
		}
		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil

}
