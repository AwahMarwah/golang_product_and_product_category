package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"

	productController "golang_product_and_product_category/controller/product"
	"golang_product_and_product_category/database"
	productModel "golang_product_and_product_category/model/product"
	productCategoryModel "golang_product_and_product_category/model/product_category"
	"golang_product_and_product_category/router"
)

// Helper function to get a random product category ID from the database
func GetRandomCategoryID(db *gorm.DB) (uuid string, err error) {
	var category productCategoryModel.ProductCategory
	err = db.Model(&productCategoryModel.ProductCategory{}).Select("id").Order("RANDOM()").First(&category).Error
	if err != nil {
		return "", err
	}
	return category.Id, nil
}

// Helper function to create a sample product in the database for unit test api list products
func createSampleProduct(db *gorm.DB, categoryId string) error {
	product := productModel.Product{
		Name:        "sample product for list",
		Description: "sample description for list",
		CategoryId:  categoryId,
	}
	return db.Create(&product).Error
}

// Set environment variable for DSN and clean up after tests
func setupEnv() func() {
	os.Setenv("DATABASE_DSN", "postgresql://asepsaepul:password@127.0.0.1:5432/product?sslmode=disable")
	return func() {
		os.Unsetenv("DATABASE_DSN")
	}
}

// TestDatabaseOpen checks if the database connection can be established
func TestDatabaseOpen(t *testing.T) {
	teardown := setupEnv()
	defer teardown()

	// Open the database connection
	db, err := database.Open()

	// Ensure there are no errors in establishing a connection
	assert.NoError(t, err)
	assert.NotNil(t, db.GormDb)
	assert.NotNil(t, db.SqlDb)

	// Close the connection after the test
	defer db.SqlDb.Close()
}

// TestRouterRun checks if the router can be initialized and run without errors
func TestRouterRun(t *testing.T) {
	teardown := setupEnv()
	defer teardown()

	// Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Open the database connection
	db, err := database.Open()
	assert.NoError(t, err)

	// Run the router setup
	err = router.Run(db)

	// Ensure the router runs without error
	assert.NoError(t, err)
}

// TestProductControllerCreate checks if the product controller's Create method works correctly
func TestProductControllerCreate(t *testing.T) {
	teardown := setupEnv()
	defer teardown()

	// Open the database connection
	db, err := database.Open()
	assert.NoError(t, err)

	// Retrieve a random CategoryId from the database
	categoryId, err := GetRandomCategoryID(db.GormDb)
	assert.NoError(t, err, "Failed to get random CategoryId")

	// Initialize the controller
	productController := productController.NewController(db.GormDb)

	// Create a Gin context for testing
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Create a sample request body for creating a product
	reqBody := productModel.CreateReqBody{
		Name:        "sample product",
		Description: "sample description",
		CategoryId:  categoryId,
	}

	// Convert the request body to JSON
	reqBodyJSON, err := json.Marshal(reqBody)
	assert.NoError(t, err)

	// Set up the request and response recorder
	ctx.Request = httptest.NewRequest("POST", "/product", bytes.NewBuffer(reqBodyJSON))
	ctx.Request.Header.Set("Content-Type", "application/json")

	// Execute the controller's Create method
	productController.Create(ctx)

	// Ensure the response code is 200 OK (or the expected status)
	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}

// TestProductControllerList checks if the product controller's List method works correctly
func TestProductControllerList(t *testing.T) {
	teardown := setupEnv()
	defer teardown()

	// Seed the random number generator
	rand.Seed(uint64(time.Now().UnixNano()))

	// Open the database connection
	db, err := database.Open()
	assert.NoError(t, err)

	// Retrieve a random CategoryId from the database
	categoryId, err := GetRandomCategoryID(db.GormDb)
	assert.NoError(t, err, "Failed to get random CategoryId")

	// Create a sample product for testing
	err = createSampleProduct(db.GormDb, categoryId)
	assert.NoError(t, err, "Failed to create sample product")

	// Generate random values for page and limit between 1 and 10
	page := rand.Intn(10) + 1
	limit := rand.Intn(10) + 1

	// Initialize the controller
	productController := productController.NewController(db.GormDb)

	// Create a Gin context for testing
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())

	// Set up the request with query parameters for page and limit
	query := "?page=" + strconv.Itoa(page) + "&limit=" + strconv.Itoa(limit)

	// Set up the request and response recorder
	ctx.Request = httptest.NewRequest("GET", "/product"+query, nil)

	// Execute the controller's List method
	productController.List(ctx)

	// Ensure the response code is 200 OK (or the expected status)
	assert.Equal(t, http.StatusOK, ctx.Writer.Status())
}
