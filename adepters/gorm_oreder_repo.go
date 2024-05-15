package adepters

import (
	"github.com/JimmyTanapon/Go-Clean-Architecture/entities"
	"github.com/JimmyTanapon/Go-Clean-Architecture/usecase"
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) usecase.OrderRepository {
	return &GormRepository{db: db}
}
func (r *GormRepository) Save(order entities.Order) error {
	if result := r.db.Create(&order); result != nil {
		return result.Error
	}
	return nil
}
