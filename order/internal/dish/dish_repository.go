package dish

import (
	"context"
	"database/sql"
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

func (r *repository) GetAll(ctx context.Context) ([]Dish, error) {
	dishes := make([]Dish, 0)
	query := "SELECT * FROM dish"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return dishes, err
	}
	for rows.Next() {
		var dish Dish
		err = rows.Scan(&dish.Id, &dish.Name, &dish.Description, &dish.Price, &dish.Quantity, &dish.IsAvailable, &dish.CreatedAt, &dish.UpdatedAt)
		if err != nil {
			return dishes, err
		}
		dishes = append(dishes, dish)
	}
	return dishes, nil
}

func (r *repository) CreateDish(ctx context.Context, dish *CreateDishReq) (*Dish, error) {
	d := Dish{
		Name:        dish.Name,
		Description: dish.Description,
		Price:       dish.Price,
		Quantity:    dish.Quantity,
		IsAvailable: dish.IsAvailable,
	}
	var lastInsertId int
	query := "INSERT INTO \"dish\"(name, description, price, quantity, is_available) VALUES ($1, $2, $3, $4, $5) returning id, created_at, updated_at"
	err := r.db.QueryRowContext(ctx, query, dish.Name, dish.Description, dish.Price, dish.Quantity, dish.IsAvailable).Scan(&lastInsertId, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return &Dish{}, err
	}

	d.Id = int64(lastInsertId)
	return &d, nil
}

func (r *repository) GetDish(ctx context.Context, id int64) (*Dish, error) {
	d := Dish{}
	query := "SELECT id, name, description, price, quantity, is_available, created_at, updated_at FROM \"dish\" WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&d.Id, &d.Name, &d.Description, &d.Price, &d.Quantity, &d.IsAvailable, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return &Dish{}, nil
	}

	return &d, nil
}

func (r *repository) UpdateDish(ctx context.Context, dish *UpdateDishReq) (*Dish, error) {
	d := Dish{
		Id:          dish.Id,
		Name:        dish.Name,
		Description: dish.Description,
		Price:       dish.Price,
		Quantity:    dish.Quantity,
		IsAvailable: dish.IsAvailable,
	}
	query := "UPDATE dish SET name = $1, description = $2, price = $3, quantity = $4, is_available = $5 WHERE id = $6 RETURNING created_at, updated_at"
	err := r.db.QueryRowContext(ctx, query, dish.Name, dish.Description, dish.Price, dish.Quantity, dish.IsAvailable, dish.Id).Scan(&d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return &Dish{}, nil
	}

	return &d, nil
}

func (r *repository) DeleteDish(ctx context.Context, id int64) (*Dish, error) {
	d := Dish{
		Id: id,
	}
	query := "DELETE FROM dish WHERE id = $1 RETURNING name, description, price, quantity, is_available, created_at, updated_at"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&d.Name, &d.Description, &d.Price, &d.Quantity, &d.IsAvailable, &d.CreatedAt, &d.UpdatedAt)
	if err != nil {
		return &Dish{}, nil
	}

	return &d, nil
}
