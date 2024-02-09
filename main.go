package main

import (
	"embed"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/template/html/v2"
	"motivation-api/delivery/http"
	"motivation-api/domain"
	"motivation-api/shared"
	httpLib "net/http"
)

//go:embed resources/*
var resourcesfs embed.FS

func main() {
	validate := validator.New()
	config, err := shared.LoadConfig(".")
	domains := domain.ConstructDomain(config.DBHost, validate)

	engine := html.NewFileSystem(httpLib.FS(resourcesfs), ".html")
	app := http.NewHttpDelivery(domains, engine)
	err = app.Listen(fmt.Sprintf(":%s", config.PortApp))
	if err != nil {
		return
	}
}
