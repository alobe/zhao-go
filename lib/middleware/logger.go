package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	z "go.uber.org/zap"
)

func Byte2String(bytes []byte) string {
	return fmt.Sprintf("%x", bytes)
}

func Logger(logger *z.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		uri := c.Request().URI()
		path := Byte2String(uri.Path())
		query := Byte2String(uri.QueryString())
		c.Next()
		cost := time.Since(start)
		logger.Info(path,
			z.Int("status", c.Response().StatusCode()),
			z.String("methods", Byte2String(c.Request().Header.Method())),
			z.String("path", path),
			z.String("query", query),
			z.String("ip", c.IP()),
			z.String("user-agent", Byte2String(c.Request().Header.UserAgent())),
			z.Duration("cost", cost),
			z.Error(c.Context().Err()),
		)
		return nil
	}
}
