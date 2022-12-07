package input

type ProductCreateInput struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required,min=0,max=50000000"`
	Stock      int    `json:"stock" binding:"required,min=5"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type ProductUpdateInput struct {
	Title      string `json:"title" binding:"required"`
	Price      int    `json:"price" binding:"required,min=0,max=50000000"`
	Stock      int    `json:"stock" binding:"required,min=5"`
	CategoryID int    `json:"category_id" binding:"required"`
}

type ProductIdUri struct {
	ID int `uri:"productId" binding:"required"`
}
