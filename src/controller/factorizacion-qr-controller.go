package controller

import (
	"github.com/Jorge2824/api-factorizacion-qr-go-v1/src/configuration"
	"github.com/Jorge2824/api-factorizacion-qr-go-v1/src/exception"
	"github.com/Jorge2824/api-factorizacion-qr-go-v1/src/model"
	"github.com/Jorge2824/api-factorizacion-qr-go-v1/src/service"
	"github.com/gofiber/fiber/v2"
)

type FactorizacionQrController struct {
	service.FactorizacionQrService
	configuration.Config
}

func NewFactorizacionQrController(
	factorizacionQrService *service.FactorizacionQrService,
	config configuration.Config) *FactorizacionQrController {
	return &FactorizacionQrController{FactorizacionQrService: *factorizacionQrService, Config: config}
}

func (controller FactorizacionQrController) Route(app *fiber.App) {
	app.Post("/v1/api/factorizacion-qr", controller.GetQrFromMatrix)
}

func (controller FactorizacionQrController) GetQrFromMatrix(c *fiber.Ctx) error {
	var request model.MatrixModel
	err := c.BodyParser(&request)
	exception.PanicLogging(err)

	response := controller.FactorizacionQrService.GetQrFromMatrix(c.Context(), request)
	return c.Status(fiber.StatusCreated).JSON(model.GeneralResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	})
}
