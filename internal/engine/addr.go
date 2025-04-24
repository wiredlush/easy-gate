package engine

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wiredlush/easy-gate/internal/routine"
)

func getAddr(status *routine.Status, c *fiber.Ctx) string {
	if status.BehindProxy {
		forwardedFor := c.Get(fiber.HeaderXForwardedFor)
		if forwardedFor != "" {
			addresses := strings.Split(forwardedFor, ",")
			return strings.TrimSpace(addresses[0])
		}
	}

	return c.IP()
}
