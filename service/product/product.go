package product

import (
	productModel "golang_product_and_product_category/model/product"
	productRepo "golang_product_and_product_category/repository/product"
	productCategoryRepo "golang_product_and_product_category/repository/product_category"
)

type (
	IService interface {
		Create(reqBody *productModel.CreateReqBody) (err error)
		List(reqQuery *productModel.ListReqQuery) (resData []productModel.ListProductResData, count int64, err error)

	}

	service struct {
		productRepo productRepo.IRepo
		productCategoryRepo productCategoryRepo.IRepo
	}
)

func NewService(productRepo productRepo.IRepo, productCategoryRepo productCategoryRepo.IRepo) IService {
	return &service{
		productRepo: productRepo,
		productCategoryRepo: productCategoryRepo,
	}
}