package main

import (
	"github.com/Jorge2824/api-factorizacion-qr-go-v1/src/configuration"
	"github.com/Jorge2824/api-factorizacion-qr-go-v1/src/controller"
	"github.com/Jorge2824/api-factorizacion-qr-go-v1/src/exception"
	service "github.com/Jorge2824/api-factorizacion-qr-go-v1/src/service/impl"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	//setup configuration
	config := configuration.New()

	//service
	factorizacionQrService := service.NewFactorizacionQrServiceImpl()

	//controller
	productController := controller.NewFactorizacionQrController(&factorizacionQrService, config)

	//setup fiber
	app := fiber.New(configuration.NewFiberConfiguration())
	app.Use(recover.New())
	app.Use(cors.New())

	//routing
	productController.Route(app)

	//start app
	err := app.Listen(config.Get("SERVER.PORT"))
	exception.PanicLogging(err)
}
