package user

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

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertId int
	query := "INSERT INTO \"user\"(username, password_hash, email, role) VALUES ($1, $2, $3, $4) returning id, created_at, updated_at"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email, user.Role).Scan(&lastInsertId, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return &User{}, err
	}

	user.Id = int64(lastInsertId)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query := "SELECT id, email, username, password_hash, role, created_at, updated_at FROM \"user\" WHERE email = $1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.Id, &u.Email, &u.Username, &u.Password, &u.Role, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return &User{}, nil
	}

	return &u, nil
}

func (r *repository) CreateSession(ctx context.Context, session *Session) error {
	var id int64
	query := "INSERT INTO session (user_id, session_token, expires_at) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRowContext(ctx, query, &session.UserId, &session.SessionToken, &session.ExpiresAt).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
