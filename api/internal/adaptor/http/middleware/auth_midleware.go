package middleware

import (
	"github/morng-dev/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthMidleware struct {
	jwtsecret string
}

func NewAuthMidleware(jwtsecret string) AuthMidleware {
	return AuthMidleware{jwtsecret: jwtsecret}
}

func (m *AuthMidleware) AuthRequire() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"Message": "missing access token",
			})
		}
		tokenString := strings.TrimPrefix("Bearer ", authHeader)
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"Message": "invalid or expired token",
			})
		}
		userID, err := uuid.Parse(claims.UserID)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"Message": "invalid user ID in token",
			},
			)
		}
		c.Locals("email", claims.Email)
		c.Locals("userID", userID)
		return c.Next()
	}
}
