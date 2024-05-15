package usecase

import (
	"github.com/JimmyTanapon/Go-Clean-Architecture/entities"
)

type OrderRepository interface {
	Save(order entities.Order) error
}
