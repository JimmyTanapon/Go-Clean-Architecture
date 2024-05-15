package usecase

import (
	"errors"

	"github.com/JimmyTanapon/Go-Clean-Architecture/entities"
)

type OrderUseCase interface {
	CreatOrder(order entities.Order) error
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderUseCase {
	return &OrderService{repo: repo}
}
func (s *OrderService) CreatOrder(order entities.Order) error {
	if order.Total < 0 {
		return errors.New("Total must be Poseitive ")
	}
	if err := s.repo.Save(order); err != nil {
		return err
	}
	return nil
}
