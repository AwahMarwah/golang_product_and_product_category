package productCategory

import productCategoryModel "golang_product_and_product_category/model/product_category"

func (r *repo) List() (resData []productCategoryModel.ListProductCategory, err error) {
	resData = make([]productCategoryModel.ListProductCategory, 0)
	return resData, r.db.Model(&productCategoryModel.ProductCategory{}).Scan(&resData).Error
}