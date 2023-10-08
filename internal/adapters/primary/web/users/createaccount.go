package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/samverrall/hex-structure/internal/core/services/users"
)

func CreateAccount(usersAPI users.API) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.Context()

		var input struct {
			Username string `json:"username"`
		}
		if err := c.BodyParser(&input); err != nil {
			return c.SendStatus(400)
		}

		resp, err := usersAPI.CreateAccount(ctx, users.CreateAccountReq(input))
		if err != nil {
			return c.SendStatus(500)
		}

		var out struct {
			UserID string `json:"userId"`
		}
		out.UserID = resp.UserID
		return c.JSON(out)
	}
}
