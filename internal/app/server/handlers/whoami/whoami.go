package whoami

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

var WireSet = wire.NewSet(wire.Struct(new(Whoami), "*"))

type Whoami struct{}

func (w *Whoami) Whoami(c *fiber.Ctx) error {
	logrus.WithContext(c.Context()).Infof("hello from highlight.io")
	c.JSON(fiber.Map{
		"message": "Hello World!",
		"version": "0.0.1",
		"time":    time.Now().Format(time.RFC3339),
	})

	return nil
}
