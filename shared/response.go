package shared

import "motivation-api/delivery/http/dto/response"

func ConstructResponseError(statusCode int, errorMsg string) (response.ResponseErrorDTO, int) {
	resp := response.ResponseErrorDTO{
		Code:   statusCode,
		Status: errorMsg,
	}

	return resp, statusCode
}
