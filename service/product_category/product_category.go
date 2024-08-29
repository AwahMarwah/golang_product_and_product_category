package productCategory

import (
	productCategoryModel "golang_product_and_product_category/model/product_category"
	productCategoryRepo "golang_product_and_product_category/repository/product_category"
)

type (
	IService interface {
		Seed(productCategory *productCategoryModel.ProductCategory)(err error)
	}

	service struct {
		productCategoryRepo productCategoryRepo.IRepo
	}
)

func NewService(productCategoryRepo productCategoryRepo.IRepo) IService {
	return &service{productCategoryRepo: productCategoryRepo}
}