package product

import (
	"golang_product_and_product_category/lib/pagination.go"
	"golang_product_and_product_category/lib/response"
	"golang_product_and_product_category/model/product"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (c *controller) List(ctx *gin.Context) {
	var reqQuery product.ListReqQuery
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}
	reqQuery.Offset = pagination.Offset(&reqQuery.Limit, &reqQuery.Page)
	resData, count, err := c.productService.List(&reqQuery) 
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessWithPage(ctx, http.StatusOK, "", resData, reqQuery.Page, reqQuery.Limit, count)
}