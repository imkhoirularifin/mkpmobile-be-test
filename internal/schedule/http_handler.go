package schedule

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/internal/domain/interfaces"
	"go-fiber-template/lib/middleware"
	"go-fiber-template/lib/utils"

	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	scheduleService interfaces.ScheduleService
}

func NewHttpHandler(r fiber.Router, scheduleService interfaces.ScheduleService) {
	handler := &httpHandler{
		scheduleService: scheduleService,
	}

	r.Post("/", middleware.Protected(), middleware.Validate[dto.CreateScheduleRequest](), handler.Create)
	r.Get("/", handler.FindAll)
	r.Get("/:id", handler.FindByID)
	r.Put("/:id", middleware.Protected(), middleware.Validate[dto.UpdateScheduleRequest](), handler.Update)
	r.Delete("/:id", middleware.Protected(), handler.Delete)
}

// @Summary		Create a new schedule
// @Description	Create a new movie schedule
// @Tags			Schedule
// @Accept			application/json
// @Produce		application/json
// @Security		Bearer
// @Param			request	body		dto.CreateScheduleRequest	true	"Schedule creation request"
// @Success		201		{object}	dto.ResponseDto{data=dto.ScheduleDto}
// @Failure		400		{object}	dto.ResponseDto
// @Failure		401		{object}	dto.ResponseDto
// @Failure		500		{object}	dto.ResponseDto
// @Router			/schedules [post]
func (h *httpHandler) Create(c *fiber.Ctx) error {
	req := utils.ExtractStructFromValidator[dto.CreateScheduleRequest](c)
	data, err := h.scheduleService.Create(c, req)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(dto.ResponseDto{
		Message: "Schedule created successfully",
		Data:    data,
	})
}

// @Summary		Get all schedules
// @Description	Get all movie schedules
// @Tags			Schedule
// @Accept			application/json
// @Produce		application/json
// @Success		200	{object}	dto.ResponseDto{data=[]dto.ScheduleDto}
// @Failure		500	{object}	dto.ResponseDto
// @Router			/schedules [get]
func (h *httpHandler) FindAll(c *fiber.Ctx) error {
	data, err := h.scheduleService.FindAll(c)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Schedules retrieved successfully",
		Data:    data,
	})
}

// @Summary		Get schedule by ID
// @Description	Get a movie schedule by ID
// @Tags			Schedule
// @Accept			application/json
// @Produce		application/json
// @Param			id	path		int	true	"Schedule ID"
// @Success		200	{object}	dto.ResponseDto{data=dto.ScheduleDto}
// @Failure		400	{object}	dto.ResponseDto
// @Failure		404	{object}	dto.ResponseDto
// @Failure		500	{object}	dto.ResponseDto
// @Router			/schedules/{id} [get]
func (h *httpHandler) FindByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid schedule ID")
	}

	data, err := h.scheduleService.FindByID(c, uint(id))
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Schedule retrieved successfully",
		Data:    data,
	})
}

// @Summary		Update schedule
// @Description	Update a movie schedule by ID
// @Tags			Schedule
// @Accept			application/json
// @Produce		application/json
// @Security		Bearer
// @Param			id		path		int							true	"Schedule ID"
// @Param			request	body		dto.UpdateScheduleRequest	true	"Schedule update request"
// @Success		200		{object}	dto.ResponseDto{data=dto.ScheduleDto}
// @Failure		400		{object}	dto.ResponseDto
// @Failure		401		{object}	dto.ResponseDto
// @Failure		404		{object}	dto.ResponseDto
// @Failure		500		{object}	dto.ResponseDto
// @Router			/schedules/{id} [put]
func (h *httpHandler) Update(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid schedule ID")
	}

	req := utils.ExtractStructFromValidator[dto.UpdateScheduleRequest](c)
	data, err := h.scheduleService.Update(c, uint(id), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Schedule updated successfully",
		Data:    data,
	})
}

// @Summary		Delete schedule
// @Description	Delete a movie schedule by ID
// @Tags			Schedule
// @Accept			application/json
// @Produce		application/json
// @Security		Bearer
// @Param			id	path		int	true	"Schedule ID"
// @Success		200	{object}	dto.ResponseDto
// @Failure		400	{object}	dto.ResponseDto
// @Failure		401	{object}	dto.ResponseDto
// @Failure		404	{object}	dto.ResponseDto
// @Failure		500	{object}	dto.ResponseDto
// @Router			/schedules/{id} [delete]
func (h *httpHandler) Delete(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid schedule ID")
	}

	if err := h.scheduleService.Delete(c, uint(id)); err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.ResponseDto{
		Message: "Schedule deleted successfully",
	})
}
