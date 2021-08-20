package services

import (
	"strconv"
	"strings"

	"github.com/ygorrodriguesmeli/gorestapi/src/errors"
	"github.com/ygorrodriguesmeli/gorestapi/src/repositories"

	"github.com/ygorrodriguesmeli/gorestapi/src/models"
)

type productService struct{}

var (
	ProductService = productService{}
)

func (s *productService) GetProductById(id string) (models.Product, errors.ApiError) {
	id = strings.TrimSpace(id)
	productId, err := strconv.Atoi(id)
	if err != nil {
		return models.Product{}, errors.NewBadRequestApiError("invalid product id")
	}
	result, err := repositories.ProductRepository.GetById(productId)
	if err != nil {
		return models.Product{}, errors.NewNotFoundApiError("cannot find product")
	}
	return result, nil
}
