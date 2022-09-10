package middleware

import (
	"time"
	"zhao-go/lib/util"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func Logger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		uri := c.Request().URI()
		path := util.String(uri.Path())
		query := util.String(uri.QueryString())
		c.Next()
		cost := time.Since(start)
		log.Info().
			Int("status", c.Response().StatusCode()).
			Str("methods", util.String(c.Request().Header.Method())).
			Str("path", path).
			Str("query", query).
			Str("ip", c.IP()).
			Str("ua", util.String(c.Request().Header.UserAgent())).
			Dur("cost", cost).
			Err(c.Context().Err()).
			Msg("router middleware inspact")
		return nil
	}
}
