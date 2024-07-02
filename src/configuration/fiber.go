package configuration

import (
	"github.com/Jorge2824/api-factorizacion-qr-go-v1/src/exception"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
