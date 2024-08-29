package product

type (
	CreateReqBody struct {
		Name           string `binding:"required"`
		Description    string `binding:"required"`
		CategoryId     string `binding:"required" json:"cage_category_id"`
	}
)