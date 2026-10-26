package grpcserver

import (
	"context"

	"github.com/AmIAFreeMan/Go-HR-managment-exercise/internal/usecase"
	hrv1 "github.com/AmIAFreeMan/Go-HR-managment-exercise/proto/hr/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmployeeServiceServer struct {
	hrv1.UnimplementedEmployeeServiceServer
	usecase *usecase.EmployeeUsecase
}

func NewEmployeeServiceServer(uc *usecase.EmployeeUsecase) *EmployeeServiceServer {
	return &EmployeeServiceServer{usecase: uc}
}

func (s *EmployeeServiceServer) CreateEmployee(ctx context.Context, req *hrv1.CreateEmployeeRequest) (*hrv1.EmployeeResponse, error) {
	emp, err := s.usecase.CreateEmployee(ctx, req.GetFirstName(), req.GetLastName(), req.GetEmail(), req.GetDepartment())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create employee: %v", err)
	}

	return &hrv1.EmployeeResponse{
		Employee: &hrv1.Employee{
			Id:         emp.ID,
			FirstName:  emp.FirstName,
			LastName:   emp.LastName,
			Email:      emp.Email,
			Department: emp.Department,
		},
	}, nil
}

func (s *EmployeeServiceServer) GetEmployee(ctx context.Context, req *hrv1.GetEmployeeRequest) (*hrv1.EmployeeResponse, error) {
	emp, err := s.usecase.GetEmployee(ctx, req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "employee not found: %v", err)
	}

	return &hrv1.EmployeeResponse{
		Employee: &hrv1.Employee{
			Id:         emp.ID,
			FirstName:  emp.FirstName,
			LastName:   emp.LastName,
			Email:      emp.Email,
			Department: emp.Department,
		},
	}, nil
}

func (s *EmployeeServiceServer) UpdateEmployee(ctx context.Context, req *hrv1.UpdateEmployeeRequest) (*hrv1.EmployeeResponse, error) {
	emp, err := s.usecase.UpdateEmployee(ctx, req.GetId(), req.GetFirstName(), req.GetLastName(), req.GetEmail(), req.GetDepartment())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to update employee: %v", err)
	}

	return &hrv1.EmployeeResponse{
		Employee: &hrv1.Employee{
			Id:         emp.ID,
			FirstName:  emp.FirstName,
			LastName:   emp.LastName,
			Email:      emp.Email,
			Department: emp.Department,
		},
	}, nil
}

func (s *EmployeeServiceServer) DeleteEmployee(ctx context.Context, req *hrv1.DeleteEmployeeRequest) (*hrv1.DeleteEmployeeResponse, error) {
	if err := s.usecase.DeleteEmployee(ctx, req.GetId()); err != nil {
		return nil, status.Errorf(codes.NotFound, "failed to delete employee: %v", err)
	}

	return &hrv1.DeleteEmployeeResponse{Success: true}, nil
}
