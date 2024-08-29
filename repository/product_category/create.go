package productCategory

import productCategoryModel "golang_product_and_product_category/model/product_category"


func (r *repo) Create(productCategory *productCategoryModel.ProductCategory) (err error) {
	return r.db.Create(productCategory).Error
}