package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

func KeyAuth(keys []string) fiber.Handler {
	return keyauth.New(keyauth.Config{
		KeyLookup: "header:X-Api-Key",
		Validator: func(c *fiber.Ctx, key string) (bool, error) {
			for _, apiKey := range keys {
				if apiKey == key {
					return true, nil
				}
			}
			return false, nil
		},
	})
}
