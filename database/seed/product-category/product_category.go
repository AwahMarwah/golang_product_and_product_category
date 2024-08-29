package main

import (
	"fmt"
	"golang_product_and_product_category/common"
	"golang_product_and_product_category/database"
	productCategoryModel "golang_product_and_product_category/model/product_category"
	productCategoryRepo "golang_product_and_product_category/repository/product_category"
	productCategory "golang_product_and_product_category/service/product_category"
	"log"
	"strings"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var name string
	fmt.Print("name: ")
	if _, err := fmt.Scanln(&name); err != nil && err.Error() != common.UnexpectedNewline {
		log.Fatal(err)
	}
	
	req := productCategoryModel.ProductCategory{
		Name: strings.ToLower(name),
	}
	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		log.Fatal(err)
	}
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.SqlDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	productCategoryService := productCategory.NewService(productCategoryRepo.NewRepo(db.GormDb))
	log.Println(req, "seed")
	if err = productCategoryService.Seed(&req); err != nil {
		log.Fatal(err)
	}
	log.Print(common.SuccessfullyCreated)
}