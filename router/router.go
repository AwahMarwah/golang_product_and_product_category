package router

import (
	"golang_product_and_product_category/controller/health"
	"golang_product_and_product_category/controller/product"
	productCategory "golang_product_and_product_category/controller/product_category"
	"golang_product_and_product_category/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(db database.DB) (err error) {
	router := gin.Default()
	router.Use(cors.New(corsConfig))
	healthController := health.NewController(db.SqlDb)
	router.GET("health", healthController.Check)
	productGroup := router.Group("product") 
	{
		productController := product.NewController(db.GormDb)
		productGroup.POST("", productController.Create)
		productGroup.GET("", productController.List)
		
	}
	productCategoryController := productCategory.NewController(db.GormDb)
	router.GET("product-category", productCategoryController.List)
	return router.Run()
}