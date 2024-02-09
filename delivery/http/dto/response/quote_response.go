package response

type ResponseErrorDTO struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}
