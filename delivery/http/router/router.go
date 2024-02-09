package router

import (
	"github.com/gofiber/fiber/v2"
	"motivation-api/delivery/http/controller"
	"motivation-api/domain"
)

func NewRouter(app *fiber.App, domain domain.Domain) {
	quoteController := controller.NewQuoteController(domain)

	app.Post("/", quoteController.CreateQuote)
	app.Get("/", quoteController.GetQuote)
}
