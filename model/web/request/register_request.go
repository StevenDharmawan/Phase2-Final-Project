package request

type RegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}
