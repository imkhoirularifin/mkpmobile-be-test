package schedule

import (
	"go-fiber-template/internal/domain/entity"
	"go-fiber-template/internal/domain/interfaces"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(data *entity.Schedule) error {
	return r.db.Create(data).Error
}

func (r *repository) FindByID(id uint) (*entity.Schedule, error) {
	var schedule entity.Schedule
	if err := r.db.First(&schedule, id).Error; err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *repository) FindAll() ([]entity.Schedule, error) {
	var schedules []entity.Schedule
	if err := r.db.Find(&schedules).Error; err != nil {
		return nil, err
	}
	return schedules, nil
}

func (r *repository) Update(data *entity.Schedule) error {
	return r.db.Save(data).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&entity.Schedule{}, id).Error
}

func NewRepository(db *gorm.DB) interfaces.ScheduleRepository {
	return &repository{db: db}
}
