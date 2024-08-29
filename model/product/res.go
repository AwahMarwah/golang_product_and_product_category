package product

import "time"

type (
	ListProductResData struct {
		Id          string    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CategoryId  string    `json:"category_id"`
		CreatedAt   time.Time `json:"created_at"`
	}
)