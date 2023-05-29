package order

import (
	"context"
	"time"
)

type Order struct {
	Id              int64     `json:"id"`
	UserId          int64     `json:"user_id"`
	Status          string    `json:"status"`
	SpecialRequests string    `json:"special_requests"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
type OrderDish struct {
	Id       int64   `json:"id"`
	OrderId  int64   `json:"order_id"`
	DishId   int64   `json:"dish_id"`
	Quantity int64   `json:"quantity"`
	Price    float64 `json:"price"`
}

type CreateOrderReq struct {
	UserId          int64   `json:"user_id"`
	SpecialRequests string  `json:"special_requests"`
	DishIds         []int64 `json:"dish_ids"`
	Quantities      []int64 `json:"quantities"`
}

type UpdateStatusReq struct {
	Id     int64  `json:"id"`
	Status string `json:"status"`
}

type CreateOrderDishReq struct {
	OrderId  int64   `json:"order_id"`
	DishId   int64   `json:"dish_id"`
	Quantity int64   `json:"quantity"`
	Price    float64 `json:"price"`
}

type Repository interface {
	CreateOrder(ctx context.Context, req *CreateOrderReq) (*Order, error)
	GetOrder(ctx context.Context, id int64) (*Order, error)
	CreateOrderDish(ctx context.Context, req *CreateOrderDishReq) (*OrderDish, error)
	SetQuantity(ctx context.Context, id int64, quantity int64) (int64, error)
	GetQuantity(ctx context.Context, id int64) (int64, error)
	GetPrice(ctx context.Context, id int64) (float64, error)
	SetStatus(req *UpdateStatusReq)
}

type Service interface {
	Order(ctx context.Context, req *CreateOrderReq) (*Order, error)
	GetOrder(ctx context.Context, id int64) (*Order, error)
	SetStatus(req *UpdateStatusReq)
}
