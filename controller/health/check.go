package health

import (
	"golang_product_and_product_category/common"
	"golang_product_and_product_category/lib/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller) Check(ctx *gin.Context) {
	if err := c.healthService.Check(); err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(ctx, http.StatusOK, common.SuccessfullyChecked, nil)

	return
}