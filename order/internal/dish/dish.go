package dish

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type Dish struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int64     `json:"quantity"`
	IsAvailable bool      `json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateDishReq struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	IsAvailable bool    `json:"is_available"`
}

type UpdateDishReq struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
	IsAvailable bool    `json:"is_available"`
}

type Repository interface {
	CreateDish(ctx context.Context, dish *CreateDishReq) (*Dish, error)
	GetDish(ctx context.Context, id int64) (*Dish, error)
	UpdateDish(ctx context.Context, dish *UpdateDishReq) (*Dish, error)
	DeleteDish(ctx context.Context, id int64) (*Dish, error)
	GetAll(ctx context.Context) ([]Dish, error)
}

type Service interface {
	CreateDish(c context.Context, req *CreateDishReq) (*Dish, error)
	GetDish(c context.Context, id int64) (*Dish, error)
	UpdateDish(c context.Context, req *UpdateDishReq) (*Dish, error)
	DeleteDish(c context.Context, id int64) (*Dish, error)
	GetAll(c context.Context) ([]Dish, error)
	CheckRole(role string) gin.HandlerFunc
	UseId() gin.HandlerFunc
}
