package interfaces

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/entity"

	"github.com/gofiber/fiber/v2"
)

type ScheduleRepository interface {
	Create(data *entity.Schedule) error
	FindByID(id uint) (*entity.Schedule, error)
	FindAll() ([]entity.Schedule, error)
	Update(data *entity.Schedule) error
	Delete(id uint) error
}

type ScheduleService interface {
	Create(c *fiber.Ctx, req *dto.CreateScheduleRequest) (*dto.ScheduleDto, error)
	FindByID(c *fiber.Ctx, id uint) (*dto.ScheduleDto, error)
	FindAll(c *fiber.Ctx) ([]dto.ScheduleDto, error)
	Update(c *fiber.Ctx, id uint, req *dto.UpdateScheduleRequest) (*dto.ScheduleDto, error)
	Delete(c *fiber.Ctx, id uint) error
}
