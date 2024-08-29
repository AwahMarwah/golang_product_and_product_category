package productCategory

import productCategoryModel "golang_product_and_product_category/model/product_category"

func (s *service) List() (resData []productCategoryModel.ListProductCategory, err error) {
	return s.productCategoryRepo.List()
}