package product

import productModel "golang_product_and_product_category/model/product"

func (s *service) List(reqQuery *productModel.ListReqQuery) (resData []productModel.ListProductResData, count int64, err error) {
	return s.productRepo.List(reqQuery)
}