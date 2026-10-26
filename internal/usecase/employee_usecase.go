package usecase

import (
	"context"
	"errors"

	"github.com/AmIAFreeMan/Go-HR-managment-exercise/internal/domain"
	"github.com/google/uuid"
)

// Repository interface
type EmployeeRepository interface {
	Create(ctx context.Context, e *domain.Employee) error
	GetByID(ctx context.Context, id string) (*domain.Employee, error)
	Update(ctx context.Context, e *domain.Employee) error
	Delete(ctx context.Context, id string) error
}

type EmployeeUsecase struct {
	repo EmployeeRepository
}

func NewEmployeeUsecase(repo EmployeeRepository) *EmployeeUsecase {
	return &EmployeeUsecase{repo: repo}
}

func (uc *EmployeeUsecase) CreateEmployee(ctx context.Context, firstName, lastName, email, department string) (*domain.Employee, error) {

	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	emp := &domain.Employee{
		ID:         uuid.New().String(),
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Department: department,
	}

	if err := uc.repo.Create(ctx, emp); err != nil {
		return nil, err
	}

	return emp, nil
}

func (uc *EmployeeUsecase) GetEmployee(ctx context.Context, id string) (*domain.Employee, error) {
	if id == "" {
		return nil, errors.New("employee id is required")
	}
	return uc.repo.GetByID(ctx, id)
}

func (uc *EmployeeUsecase) UpdateEmployee(ctx context.Context, id, firstName, lastName, email, department string) (*domain.Employee, error) {
	emp := &domain.Employee{
		ID:         id,
		FirstName:  firstName,
		LastName:   lastName,
		Email:      email,
		Department: department,
	}

	if err := uc.repo.Update(ctx, emp); err != nil {
		return nil, err
	}

	return emp, nil
}

func (uc *EmployeeUsecase) DeleteEmployee(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("employee id is required")
	}
	return uc.repo.Delete(ctx, id)
}
