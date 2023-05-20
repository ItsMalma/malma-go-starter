package controller

import (
	"log"
	"net/http"

	"github.com/ItsMalma/malma-go-starter/exception"
	"github.com/ItsMalma/malma-go-starter/model"
	"github.com/gofiber/fiber/v2"
)

func NewErrorHandler() fiber.ErrorHandler {
	repositoryErrorStatusesCode := map[exception.RepositoryErrorStatus]int{
		exception.RepositoryErrorStatusNotFound: 404,
	}

	return func(c *fiber.Ctx, err error) error {
		switch err := err.(type) {
		case *fiber.Error:
			return c.Status(err.Code).JSON(model.Payload{
				Status: model.ToPayloadStatus(http.StatusText(err.Code)),
				Error:  err.Message,
			})
		case exception.ControllerError:
			return c.Status(err.Code).JSON(model.Payload{
				Status: err.Status,
				Data:   nil,
				Error:  err.Message,
			})
		case exception.ValidatorErrors:
			return c.Status(400).JSON(model.Payload{
				Status: "BAD_REQUEST",
				Data:   nil,
				Error:  err,
			})
		case exception.ValidatorError:
			return c.Status(400).JSON(model.Payload{
				Status: "BAD_REQUEST",
				Data:   nil,
				Error:  err,
			})
		case exception.RepositoryError:
			return c.Status(repositoryErrorStatusesCode[err.Status]).JSON(model.Payload{
				Status: string(err.Status),
				Data:   nil,
				Error:  err.Message,
			})
		default:
			log.Println(err)
			return c.Status(500).JSON(model.Payload{
				Status: "INTERNAL_SERVER_ERROR",
				Error:  "Internal Server Error",
			})
		}
	}
}
