package request

import "motivation-api/domain/entity"

type CreateQuoteRequest struct {
	Quote string `json:"quote"`
}

func (u CreateQuoteRequest) ToCreateQuote() entity.Quote {
	return entity.Quote{
		Quote: u.Quote,
	}
}
