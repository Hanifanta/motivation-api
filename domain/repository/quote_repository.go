package repository

import (
	"gorm.io/gorm"
	"log"
)

type QuoteRepository interface {
	CreateQuote(data any) error
	FindAll(data any, conditions ...any) error
}

type QuoteRepositoryImpl struct {
	database *gorm.DB
}

func NewQuoteRepository(database *gorm.DB) QuoteRepositoryImpl {
	return QuoteRepositoryImpl{database: database}
}

func (q QuoteRepositoryImpl) CreateQuote(data any) error {
	res := q.database.Create(data)
	if res.Error != nil {
		log.Printf("Error while creating quote: %v", res.Error)
		return res.Error
	}
	return nil
}

func (q QuoteRepositoryImpl) FindAll(data any, conditions ...any) error {
	result := q.database.Order("RAND()").Limit(1).Find(data, conditions...)
	if result.Error != nil {
		log.Printf("Error while finding all quotes: %v", result.Error)
		return result.Error
	}
	return nil
}
