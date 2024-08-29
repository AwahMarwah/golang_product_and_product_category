package health

import (
	"database/sql"
	healthRepo"golang_product_and_product_category/repository/health"
	"golang_product_and_product_category/service/health"
)

type controller struct {
	healthService health.IService
}

func NewController(db *sql.DB) *controller {
	return &controller{healthService: health.NewService(healthRepo.NewRepo(db))}
}