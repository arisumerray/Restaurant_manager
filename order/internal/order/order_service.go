package order

import (
	"context"
	"log"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) Order(ctx context.Context, req *CreateOrderReq) (*Order, error) {
	order, err := s.CreateOrder(ctx, req)
	if err != nil {
		log.Print(order)
		return order, err
	}
	for index, id := range req.DishIds {
		price, err := s.Repository.GetPrice(ctx, id)
		if err != nil {
			return order, err
		}

		_, err = s.Repository.CreateOrderDish(ctx, &CreateOrderDishReq{
			OrderId:  order.Id,
			DishId:   id,
			Quantity: req.Quantities[index],
			Price:    price,
		})
		if err != nil {
			return order, err
		}
	}

	return order, nil
}

func (s *service) SetStatus(req *UpdateStatusReq) {
	time.Sleep(time.Second * 5)
	s.Repository.SetStatus(req)
}

func (s *service) GetOrder(ctx context.Context, id int64) (*Order, error) {
	order, err := s.Repository.GetOrder(ctx, id)
	return order, err
}
