package infrastructure

import (
	"go-fiber-template/internal/auth"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/internal/schedule"
	"go-fiber-template/internal/user"
	"go-fiber-template/lib/config"
	"go-fiber-template/lib/database"
	"go-fiber-template/lib/xlogger"
	"go-fiber-template/lib/xvalidator"

	"gorm.io/gorm"
)

var (
	cfg        config.AppConfig
	dbInstance *database.Database
	db         *gorm.DB

	authService     interfaces.AuthService
	userService     interfaces.UserService
	scheduleService interfaces.ScheduleService
)

func init() {
	cfg = config.Setup()
	xlogger.Setup(cfg)
	xvalidator.Setup()

	dbInstance = database.New(database.Config{
		Driver: cfg.Database.Driver,
		Dsn:    cfg.Database.DbString,
	})
	db = dbInstance.GetDB()

	userRepository := user.NewRepository(db)
	scheduleRepository := schedule.NewRepository(db)

	authService = auth.NewService(userRepository)
	userService = user.NewService(userRepository)
	scheduleService = schedule.NewService(scheduleRepository)
}
