package entity

type Body struct {
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age" binding:"gte=1, lte=22"`
	Email string `json:"email" validate:"required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2,max=10"`
	Description string `json:"description" binding:"max=15"`
	URL         string `json:"url" binding:"required,url"`
}
