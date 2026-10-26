package repository

import (
	"context"
	"errors"

	"github.com/AmIAFreeMan/Go-HR-managment-exercise/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository struct {
	db *pgxpool.Pool
}

func NewEmployeeRepository(db *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(ctx context.Context, e *domain.Employee) error {
	query := `INSERT INTO employees (id, first_name, last_name, email, department) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, query, e.ID, e.FirstName, e.LastName, e.Email, e.Department)
	return err
}

func (r *EmployeeRepository) GetByID(ctx context.Context, id string) (*domain.Employee, error) {
	query := `SELECT id, first_name, last_name, email, department FROM employees WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var e domain.Employee
	err := row.Scan(&e.ID, &e.FirstName, &e.LastName, &e.Email, &e.Department)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("employee not found")
		}
		return nil, err
	}
	return &e, nil
}

func (r *EmployeeRepository) Update(ctx context.Context, e *domain.Employee) error {
	query := `UPDATE employees SET first_name = $1, last_name = $2, email = $3, department = $4 WHERE id = $5`
	cmdTag, err := r.db.Exec(ctx, query, e.FirstName, e.LastName, e.Email, e.Department, e.ID)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("employee not found")
	}
	return nil
}

func (r *EmployeeRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM employees WHERE id = $1`
	cmdTag, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmdTag.RowsAffected() == 0 {
		return errors.New("employee not found")
	}
	return nil
}
