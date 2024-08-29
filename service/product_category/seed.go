package productCategory

import (
	"errors"
	"golang_product_and_product_category/common"
	productCategoryModel "golang_product_and_product_category/model/product_category"

	"gorm.io/gorm"
)


func (s *service) Seed(productCategory *productCategoryModel.ProductCategory)(err error) {
	if _, err = s.productCategoryRepo.Take([]string{"id"}, &productCategoryModel.ProductCategory{Name: productCategory.Name}); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return
		}
	} else {
		return errors.New(common.ProductCategoryAlreadyExist)
	}
	return s.productCategoryRepo.Create(productCategory)
}