package productCategory

import productCategoryModel "golang_product_and_product_category/model/product_category"


func (r *repo) Take(selectParams []string, conditions *productCategoryModel.ProductCategory) (productCategory productCategoryModel.ProductCategory, err error) {
	return productCategory, r.db.Select(selectParams).Take(&productCategory, conditions).Error
}