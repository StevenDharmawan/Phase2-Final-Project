package helper

import (
	"net/http"
	"phase2-final-project/model/web/response"
)

func ErrBadRequest(detail any) response.ErrorResponse {
	return response.ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: "Bad Request",
		Detail:  detail,
	}
}

func ErrUnauthorized(detail any) response.ErrorResponse {
	return response.ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized Access",
		Detail:  detail,
	}
}

func ErrInternalServer(detail any) response.ErrorResponse {
	return response.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal Server Error",
		Detail:  detail,
	}
}

func ErrNotFound(detail any) response.ErrorResponse {
	return response.ErrorResponse{
		Code:    http.StatusNotFound,
		Message: "Error Not Found",
		Detail:  detail,
	}
}
