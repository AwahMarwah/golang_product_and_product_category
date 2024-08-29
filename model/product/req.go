package product

type (
	CreateReqBody struct {
		Name           string `binding:"required"`
		Description    string `binding:"required"`
		CategoryId     string `binding:"required" json:"cage_category_id"`
	}
	ListReqQuery struct {
		Limit       int    `form:"limit"`
		Offset      int
		Page        int    `form:"page"`
		Name        string `form:"name"`
		Description string `form:"description"`
		CategoryId  string `form:"category_id"`
	}
)