package controller

import (
	"github.com/gofiber/fiber/v2"
	"motivation-api/delivery/http/dto/request"
	"motivation-api/domain"
	"motivation-api/shared"
)

type QuoteController struct {
	domain domain.Domain
}

func NewQuoteController(domain domain.Domain) *QuoteController {
	return &QuoteController{domain: domain}
}

func (q *QuoteController) CreateQuote(ctx *fiber.Ctx) error {
	quote := ctx.FormValue("quote")
	if quote == "" {
		resp, statusCode := shared.ConstructResponseError(fiber.StatusBadRequest, "Invalid request")
		return ctx.Status(statusCode).JSON(resp)
	}

	quoteRequest := request.CreateQuoteRequest{Quote: quote}
	err := q.domain.QuoteUsecase.CreateQuote(quoteRequest.ToCreateQuote())
	if err != nil {
		resp, statusCode := shared.ConstructResponseError(fiber.StatusInternalServerError, err.Error())
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Redirect("/")
}

func (q *QuoteController) GetQuote(ctx *fiber.Ctx) error {
	quotes, err := q.domain.QuoteUsecase.GetAllQuotes()
	if err != nil {
		resp, statusCode := shared.ConstructResponseError(fiber.StatusInternalServerError, err.Error())
		return ctx.Status(statusCode).JSON(resp)
	}

	return ctx.Render("resources/views/home", fiber.Map{
		"Quote": quotes.Quote,
	})
}
