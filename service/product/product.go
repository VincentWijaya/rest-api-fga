package product

import (
	"rest-api/core"
	"rest-api/repo/product"
)

type Service struct {
	productRepo *product.Repository
}

func NewService(productRepo *product.Repository) *Service {
	return &Service{
		productRepo: productRepo,
	}
}

func (s *Service) GetAllProducts() ([]core.Product, error) {
	res, err := s.productRepo.GetAllProducts()

	return res, err
}
