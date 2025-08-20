package exerr

import (
	"errors"

	"github.com/gofiber/fiber/v3"
)

type HttpError struct {
	Message string `json:"message"`
}

func SendHTTP(ctx fiber.Ctx, err error) error {
	exerr, ok := to(err)
	if !ok {
		code := fiber.StatusInternalServerError
		message := "Internal server error"

		var ferr *fiber.Error
		if errors.As(err, &ferr) {
			code = ferr.Code
			message = ferr.Message
		}

		return ctx.Status(code).JSON(fiber.Map{
			"message": message,
		})
	}

	body := HttpError{
		Message: exerr.Message(),
	}

	return ctx.Status(exerr.Code()).JSON(body)
}

func SendHTTPWithStatus(ctx fiber.Ctx, err error, status int) error {
	exerr, ok := to(err)
	if !ok {

		return ctx.Status(status).JSON(HttpError{
			Message: "unknown error",
		})
	}

	body := HttpError{
		Message: exerr.Message(),
	}

	return ctx.Status(status).JSON(body)
}
