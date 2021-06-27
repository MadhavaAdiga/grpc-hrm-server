package db

import (
	"context"

	"github.com/google/uuid"
)

// Store is a base interface
// defines set of method to be implemented by different stores
// interface holding different database queries to be performed
type Store interface {
	// organizations
	CreateOrganization(ctx context.Context, param CreateOrganizationParam) (Organization, error)
	FindOrganizationByName(ctx context.Context, name string) (Organization, error)
	FindOrganizationByID(ctx context.Context, id uuid.UUID) (Organization, error)
	// users
	CreateUser(ctx context.Context, arg CreateUserParam) (User, error)
	FindUserByName(ctx context.Context, userName string) (User, error)
	FindUserById(ctx context.Context, id uuid.UUID) (User, error)
	// roles
	CreateRole(ctx context.Context, arg CreateRoleParam) (Role, error)
	FindRoleByOrganizationID(ctx context.Context, arg FindRoleByOrgIDParam) (Role, error)
	FindRoleByOrganizationName(ctx context.Context, arg FindRoleByOrgNameParam) (Role, error)
	// employees
	CreateEmployee(ctx context.Context, arg CreateEmployeeParam) (Employee, error)
	FindEmployeeByUnameAndOrg(ctx context.Context, arg FindEmployeeUnameAndOrgParam) (Employee, error)
	// payrolls
	CreatePayroll(ctx context.Context, arg CreatePayrollParam) (Payroll, error)
	FindPayrollByEmp(ctx context.Context, id uuid.UUID) (Payroll, error)
}

// A compile time check to make sure Oueries implements Querier
// var _ Store = (*SQLStore)(nil)
