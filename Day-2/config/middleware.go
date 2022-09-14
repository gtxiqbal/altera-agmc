package config

import (
	"github.com/google/uuid"
	"github.com/gtxiqbal/altera-agmc/Day-2/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"strings"
)

func CustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if strings.HasSuffix(c.Request().URL.EscapedPath(), "/") {
			c.Request().URL.RawPath = strings.TrimRight(c.Request().URL.EscapedPath(), "/")
			c.Request().URL.Path = strings.TrimRight(c.Request().URL.EscapedPath(), "/")
		}
		return next(c)
	}
}

func RecoverConfig() middleware.RecoverConfig {
	return middleware.RecoverConfig{
		StackSize:    1 << 10,
		LogLevel:     log.ERROR,
		LogErrorFunc: helper.CustomHttpErrorHandler,
	}
}

func RequestIDConfig() middleware.RequestIDConfig {
	return middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.NewString()
		},
	}
}

func LoggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${status} | ${id} | ${remote_ip} | ${host} | ${method} | ${path}\n",
	}
}
