package input

type CategoryCreateInput struct {
	Type string `json:"type" binding:"required"`
}

type CategoryPatchInput struct {
	Type string `json:"type" binding:"required"`
}

type CategoryIdUri struct {
	ID int `uri:"id" binding:"required"`
}
