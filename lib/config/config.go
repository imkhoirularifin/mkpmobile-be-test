package config

import (
	"go-fiber-template/lib/common"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog/log"
)

func FiberCfg(cfg AppConfig) fiber.Config {
	return fiber.Config{
		AppName:               cfg.AppName,
		ErrorHandler:          common.ErrorHandler,
		DisableStartupMessage: true,
	}
}

var CorsCfg = cors.Config{
	AllowOrigins:     "http://localhost:3000",
	AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
	AllowHeaders:     "*",
	AllowCredentials: true,
}

func FiberZerologCfg(cfg AppConfig) fiberzerolog.Config {
	return fiberzerolog.Config{
		Logger:          &log.Logger,
		Fields:          cfg.LogFields,
		WrapHeaders:     true,
		FieldsSnakeCase: true,
	}
}
