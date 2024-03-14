package auth

import "github.com/gofiber/fiber/v2"

type LoginStrategy interface {
	Do(ctx *fiber.Ctx) error
}

type LoginOperator struct {
	loginStrategy LoginStrategy
}

func (o *LoginOperator) SetOperator(strategy LoginStrategy) {
	o.loginStrategy = strategy
}

func (o *LoginOperator) DoLogin(ctx *fiber.Ctx) error {
	return o.loginStrategy.Do(ctx)
}
