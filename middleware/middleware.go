package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kauri646/be-idealibs/utils"
)

func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("x-token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims, err := utils.DecodeToken(token)

	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	role := claims["role"].(string)
	if role != "admin" {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "forbidden access",
		})
	}

	//ctx.Locals("userInfo", claims)

	// if token != "secret" {
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"message": "unauthenticated",
	// 	})
	// }

	return ctx.Next()
}

func PermissionsCreate(ctx *fiber.Ctx) error {
	return ctx.Next()
}
