package response

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  any    `json:"detail,omitempty"`
}
