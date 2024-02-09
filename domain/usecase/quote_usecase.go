package usecase

import (
	"motivation-api/domain/entity"
	"motivation-api/domain/repository"
)

type QuoteUsecase interface {
	CreateQuote(quote entity.Quote) error
	GetAllQuotes() (*entity.Quote, error)
}

type QuoteUsecaseImpl struct {
	QuoteRepository repository.QuoteRepository
}

func NewQuoteUsecase(quoteRepository repository.QuoteRepository) QuoteUsecaseImpl {
	return QuoteUsecaseImpl{QuoteRepository: quoteRepository}
}

func (q QuoteUsecaseImpl) CreateQuote(quote entity.Quote) error {
	err := q.QuoteRepository.CreateQuote(&quote)
	if err != nil {
		return err
	}
	return nil
}

func (q QuoteUsecaseImpl) GetAllQuotes() (*entity.Quote, error) {
	var quotes entity.Quote
	err := q.QuoteRepository.FindAll(&quotes)
	if err != nil {
		return nil, err
	}

	return &quotes, nil
}
