package productCategory

import (
	"golang_product_and_product_category/lib/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller) List(ctx *gin.Context) {
	resData, err := c.productCategoryService.List()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, "", resData)
}