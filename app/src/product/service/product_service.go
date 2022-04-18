package service

import "loja/app/src/product/model"

type ProductService struct {
	Persistence model.ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (model.ProductInterface, error) {
	product, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}
