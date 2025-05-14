package controller

import (
	"go-api/model"
	usecase "go-api/useCase"

	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

// gin.context encapsula tanto os dados da requisição quanto os dados da resposta
// e o contexto da requisição
func (p *productController) GetProducts(c *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, products)
}

// c === contexto pois o gin.context serve para
// encapsular a requisição HTTP atual e para encapsular tambem
// os metodos que constroem a resposta HTTP
func (p *productController) CreateProduct(c *gin.Context) {
	// como quando o metodo post é enviado ele envia o produto como
	// json, então precisamos fazer com que o json
	// vire a esrutura de product do projeto
	var product model.Product
	// função que pega o json da requisição e tranforma para
	// a estrutura de product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	insertedProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, insertedProduct)

}
