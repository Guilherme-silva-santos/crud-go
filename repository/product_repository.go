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

// cria as consultas para o banco de dados
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
func (pr *ProductRepository) FindProductsById(id_product int) (model.Product, error) {
	// query qye será feita no bano

	query, err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")

	if err != nil {
		fmt.Println("Error preparing query:", err)
		return model.Product{}, err
	}

	var productObj model.Product

	err = query.QueryRow(id_product).Scan(&productObj.ID, &productObj.Name, &productObj.Price)

	if err != nil {
		// se o banco não encontou nenhum registro
		if err == sql.ErrNoRows {
			return model.Product{}, nil
		}
		return model.Product{}, err
	}
	// termina a query
	query.Close()
	return productObj, nil

}

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {

	var id int

	query, err := pr.connection.Prepare("INSERT INTO product (product_name, price) VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return 0, err
	}

	//parametros que devem ser passados para a query
	// o retorno da query é o id do produto
	// o id do produto é retornado na variavel product.ID
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println("Error preparing query:", err)
		return 0, err
	}

	query.Close()
	return id, nil
}
