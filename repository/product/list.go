package product

import productModel "golang_product_and_product_category/model/product"

func (r *repo) List(reqQuery *productModel.ListReqQuery) (resData []productModel.ListProductResData, count int64, err error) {
	resData = make([]productModel.ListProductResData, 0)
	condition := ""
	if reqQuery.Name != "" {
		condition += "name ILIKE '%" + reqQuery.Name + "%'"
	}
	if reqQuery.Description != "" {
		condition += "description ILIKE '%" + reqQuery.Description + "%'"
	}
	if reqQuery.CategoryId != "" {
		condition = "category_id = '" + reqQuery.CategoryId + "'"
	}
	return resData, count, r.db.Model(&productModel.Product{}).Where(condition).Count(&count).Order("created_at DESC").Limit(reqQuery.Limit).Offset(reqQuery.Offset).Scan(&resData).Error
}
