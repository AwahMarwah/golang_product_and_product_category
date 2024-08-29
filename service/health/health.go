package health

import healthRepo"golang_product_and_product_category/repository/health"

type (
	IService interface {
		Check() (err error)
	}

	service struct {
		healthRepo healthRepo.IRepo
	}
)

func NewService(healthRepo healthRepo.IRepo) IService {
	return &service{healthRepo: healthRepo}
}
