package request

type UpdateUserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
