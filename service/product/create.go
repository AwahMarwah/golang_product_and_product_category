package product

import (
	"errors"
	"golang_product_and_product_category/common"
	productModel "golang_product_and_product_category/model/product"
	productCategoryModel "golang_product_and_product_category/model/product_category"

	"gorm.io/gorm"
)

func (s *service) Create(reqBody *productModel.CreateReqBody) (err error) {
	_, err = s.productCategoryRepo.Take([]string{"id"}, &productCategoryModel.ProductCategory{Id: reqBody.CategoryId})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New(common.ProductCategoryNotFound)
		}
		return err
	}
	return s.productRepo.Create(&productModel.Product{
		Name: reqBody.Name,
		Description: reqBody.Description,
		CategoryId: reqBody.CategoryId,
	})
}