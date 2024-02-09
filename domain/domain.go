package domain

import (
	"github.com/go-playground/validator/v10"
	"motivation-api/domain/repository"
	"motivation-api/domain/usecase"
	"motivation-api/infrastructure"
)

type Domain struct {
	QuoteUsecase usecase.QuoteUsecase
	Validate     *validator.Validate
}

func ConstructDomain(dbDsn string, validate *validator.Validate) Domain {
	databaseConnection := infrastructure.NewDatabaseConnection(dbDsn)

	quoteRepository := repository.NewQuoteRepository(databaseConnection)
	quoteUsecase := usecase.NewQuoteUsecase(quoteRepository)

	return Domain{
		QuoteUsecase: &quoteUsecase,
		Validate:     validate,
	}
}
