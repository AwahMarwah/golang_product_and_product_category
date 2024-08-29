package main

import (
	_ "github.com/joho/godotenv/autoload"
	"golang_product_and_product_category/database"
	"golang_product_and_product_category/router"
	"log"
)

func main() {
	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = db.SqlDb.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	if err = router.Run(db); err != nil {
		log.Fatal(err)
	}
}