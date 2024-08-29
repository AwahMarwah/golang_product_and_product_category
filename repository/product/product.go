package product

import (
	productModel "golang_product_and_product_category/model/product"

	"gorm.io/gorm"
)

type (
	IRepo interface {
		Create(product *productModel.Product) (err error)
		List(reqQuery *productModel.ListReqQuery) (resData []productModel.ListProductResData, count int64, err error)

	}
	
	repo struct {
		db *gorm.DB
	}
)

func NewRepo(db *gorm.DB) IRepo {
	return &repo{db: db}
}