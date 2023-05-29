package order

import (
	"context"
	"database/sql"
	"log"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateOrder(ctx context.Context, req *CreateOrderReq) (*Order, error) {
	order := Order{
		Status:          "in_progress",
		UserId:          req.UserId,
		SpecialRequests: req.SpecialRequests,
	}
	query := "INSERT INTO \"order\" (user_id, status, special_requests) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	err := r.db.QueryRowContext(ctx, query, order.UserId, order.Status, order.SpecialRequests).Scan(&order.Id, &order.CreatedAt, &order.UpdatedAt)
	return &order, err
}

func (r *repository) GetOrder(ctx context.Context, id int64) (*Order, error) {
	order := Order{}
	query := "SELECT * FROM \"order\" WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&order.Id, &order.UserId, &order.Status, &order.SpecialRequests, &order.CreatedAt, &order.UpdatedAt)
	return &order, err
}

func (r *repository) CreateOrderDish(ctx context.Context, req *CreateOrderDishReq) (*OrderDish, error) {
	orderDish := OrderDish{
		OrderId:  req.OrderId,
		DishId:   req.DishId,
		Quantity: req.Quantity,
		Price:    req.Price,
	}
	query := "INSERT INTO \"order_dish\" (order_id, dish_id, quantity, price) VALUES ($1, $2, $3, $4) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, orderDish.OrderId, orderDish.DishId, orderDish.Quantity, orderDish.Price).Scan(&orderDish.Id)
	return &orderDish, err
}

func (r *repository) GetQuantity(ctx context.Context, id int64) (int64, error) {
	var quantity int64
	query := "SELECT quantity FROM dish WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&quantity)
	if err != nil {
		return 0, err
	}
	return quantity, nil
}

func (r *repository) SetQuantity(ctx context.Context, id int64, quantity int64) (int64, error) {
	query := "UPDATE dish SET quantity = $1 WHERE id = $2 RETURNING id"
	err := r.db.QueryRowContext(ctx, query, quantity, id).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) GetPrice(ctx context.Context, id int64) (float64, error) {
	var price float64
	query := "SELECT price FROM dish WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&price)
	return price, err
}

func (r *repository) SetStatus(req *UpdateStatusReq) {
	log.Printf("try update where id = %d status = %s", req.Id, req.Status)
	//r.db.QueryRowContext(ctx, "UPDATE \"order\" SET status = $1 WHERE id = $2", req.Status, req.Id)
	db, err := sql.Open("postgres", "postgresql://root:root@localhost:5432/db?sslmode=disable")
	if err != nil {
		log.Print(err)
		return
	}
	db.QueryRow("UPDATE \"order\" SET status = $1 WHERE id = $2", req.Status, req.Id)

}
