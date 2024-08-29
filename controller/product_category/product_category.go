package productCategory

import (
	productCategoryRepo "golang_product_and_product_category/repository/product_category"
	productCategory "golang_product_and_product_category/service/product_category"

	"gorm.io/gorm"
)

type controller struct {
	productCategoryService productCategory.IService
}

func NewController(db *gorm.DB) *controller {
	return &controller{productCategoryService: productCategory.NewService(productCategoryRepo.NewRepo(db))}
}