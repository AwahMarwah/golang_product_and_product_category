package productCategory

import (
	productCategoryModel "golang_product_and_product_category/model/product_category"

	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(productCategory *productCategoryModel.ProductCategory) (err error)
		List() (resData []productCategoryModel.ListProductCategory, err error)
		Take(selectParams []string, conditions *productCategoryModel.ProductCategory) (productCategory productCategoryModel.ProductCategory, err error)
	}

	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}