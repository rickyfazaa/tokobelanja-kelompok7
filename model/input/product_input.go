package input

type ProductCreateInput struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type ProductUpdateInput struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required"`
	Stock      int    `json:"stock" binding:"required"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type ProductIDUri struct {
	ID int `uri:"id" binding:"required"`
}
