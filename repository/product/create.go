package product

import productModel "golang_product_and_product_category/model/product"

func (r *repo) Create(product *productModel.Product) (err error) {
	return r.db.Create(product).Error
}
