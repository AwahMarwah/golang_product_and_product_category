package product

import (
	productRepo "golang_product_and_product_category/repository/product"
	productCategoryRepo "golang_product_and_product_category/repository/product_category"
	"golang_product_and_product_category/service/product"

	"gorm.io/gorm"
)

type controller struct {
	productService product.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{productService: product.NewService(productRepo.NewRepo(db), productCategoryRepo.NewRepo(db))}
}