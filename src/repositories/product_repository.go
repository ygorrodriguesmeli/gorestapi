package repositories

import (
	"fmt"

	"github.com/ygorrodriguesmeli/gorestapi/src/database"
	"github.com/ygorrodriguesmeli/gorestapi/src/models"
)

type productRepository struct{}

var (
	ProductRepository = productRepository{}
)

func (r *productRepository) GetById(id int) (models.Product, error) {
	db := database.GetDatabase()
	var product models.Product
	err := db.First(&product, id).Error
	if err != nil {
		return models.Product{}, fmt.Errorf("product not found: %v", err)
	}
	return product, nil
}
