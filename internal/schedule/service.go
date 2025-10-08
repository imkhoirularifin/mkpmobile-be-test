package schedule

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/internal/domain/interfaces"
	"time"

	"github.com/gofiber/fiber/v2"
)

type service struct {
	scheduleRepo interfaces.ScheduleRepository
}

func (s *service) Create(c *fiber.Ctx, req *dto.CreateScheduleRequest) (*dto.ScheduleDto, error) {
	showDate, err := time.Parse("2006-01-02", req.ShowDate)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid show_date format")
	}

	schedule := &entity.Schedule{
		MovieTitle:     req.MovieTitle,
		StudioName:     req.StudioName,
		ShowDate:       showDate,
		ShowTime:       req.ShowTime,
		AvailableSeats: req.AvailableSeats,
		Price:          req.Price,
	}

	if err := s.scheduleRepo.Create(schedule); err != nil {
		return nil, err
	}

	return constructScheduleDto(schedule), nil
}

func (s *service) FindByID(c *fiber.Ctx, id uint) (*dto.ScheduleDto, error) {
	schedule, err := s.scheduleRepo.FindByID(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Schedule not found")
	}

	return constructScheduleDto(schedule), nil
}

func (s *service) FindAll(c *fiber.Ctx) ([]dto.ScheduleDto, error) {
	schedules, err := s.scheduleRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var scheduleDtos []dto.ScheduleDto
	for _, schedule := range schedules {
		scheduleDtos = append(scheduleDtos, *constructScheduleDto(&schedule))
	}

	return scheduleDtos, nil
}

func (s *service) Update(c *fiber.Ctx, id uint, req *dto.UpdateScheduleRequest) (*dto.ScheduleDto, error) {
	schedule, err := s.scheduleRepo.FindByID(id)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Schedule not found")
	}

	if req.MovieTitle != nil {
		schedule.MovieTitle = *req.MovieTitle
	}
	if req.StudioName != nil {
		schedule.StudioName = *req.StudioName
	}
	if req.ShowDate != nil {
		showDate, err := time.Parse("2006-01-02", *req.ShowDate)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid show_date format")
		}
		schedule.ShowDate = showDate
	}
	if req.ShowTime != nil {
		schedule.ShowTime = *req.ShowTime
	}
	if req.AvailableSeats != nil {
		schedule.AvailableSeats = *req.AvailableSeats
	}
	if req.Price != nil {
		schedule.Price = *req.Price
	}

	if err := s.scheduleRepo.Update(schedule); err != nil {
		return nil, err
	}

	return constructScheduleDto(schedule), nil
}

func (s *service) Delete(c *fiber.Ctx, id uint) error {
	schedule, err := s.scheduleRepo.FindByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Schedule not found")
	}

	if err := s.scheduleRepo.Delete(schedule.ID); err != nil {
		return err
	}

	return nil
}

func constructScheduleDto(schedule *entity.Schedule) *dto.ScheduleDto {
	return &dto.ScheduleDto{
		ID:             schedule.ID,
		MovieTitle:     schedule.MovieTitle,
		StudioName:     schedule.StudioName,
		ShowDate:       schedule.ShowDate,
		ShowTime:       schedule.ShowTime,
		AvailableSeats: schedule.AvailableSeats,
		Price:          schedule.Price,
		CreatedAt:      schedule.CreatedAt,
		UpdatedAt:      schedule.UpdatedAt,
	}
}

func NewService(scheduleRepo interfaces.ScheduleRepository) interfaces.ScheduleService {
	return &service{
		scheduleRepo: scheduleRepo,
	}
}
