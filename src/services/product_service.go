package services

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ygorrodriguesmeli/gorestapi/src/repositories"

	"github.com/ygorrodriguesmeli/gorestapi/src/models"
)

type productService struct{}

var (
	ProductService = productService{}
)

func (s *productService) GetProductById(id string) (models.Product, error) {
	id = strings.TrimSpace(id)
	productId, err := strconv.Atoi(id)
	if err != nil {
		return models.Product{}, fmt.Errorf("invalid param id: %v", err)
	}
	result, err := repositories.ProductRepository.GetById(productId)
	if err != nil {
		return models.Product{}, fmt.Errorf("cannot find product by id: %v", err)
	}
	return result, nil
}
