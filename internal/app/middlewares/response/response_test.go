package middlewares

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func Test_UnifiedResponse(t *testing.T) {
	t.Parallel()
	app := fiber.New()

	app.Use(UnifiedResponse())

	tests := []struct {
		name    string
		status  int
		message string
	}{
		{
			name:    "with status 200",
			status:  200,
			message: "OK",
		},
		{
			name:    "with status 404",
			status:  404,
			message: "Not Found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app.Get("/", func(c *fiber.Ctx) error {
				if tt.status == 200 {
					c.JSON(fiber.Map{
						"message": tt.message,
					})
					return nil
				} else {
					return fiber.NewError(tt.status, tt.message)
				}
			})

			resp, _ := app.Test(httptest.NewRequest(fiber.MethodGet, "/", nil))
			body, _ := io.ReadAll(resp.Body)

			var respData Response
			_ = json.Unmarshal(body, &respData)
			assert.Equal(t, tt.status, respData.Meta.Code)
			snaps.MatchSnapshot(t, respData.Data)
		})
	}
}
