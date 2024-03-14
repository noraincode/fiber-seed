package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Meta struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func respondWithError(c *fiber.Ctx, err error) {
	code := fiber.StatusInternalServerError
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	c.Status(code).JSON(Response{
		Meta: Meta{
			Status: http.StatusText(code),
			Code:   code,
		},
		Data: fiber.Map{
			"message": err.Error(),
		},
	})
}

func UnifiedResponse() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		if err != nil {
			respondWithError(c, err)
			return nil
		}

		originalResponse := c.Response().Body()
		var jsonBody map[string]interface{}
		_ = json.Unmarshal(originalResponse, &jsonBody)

		code := c.Response().StatusCode()
		c.JSON(Response{
			Meta: Meta{
				Status: http.StatusText(code),
				Code:   code,
			},
			Data: jsonBody,
		})

		return nil
	}
}
