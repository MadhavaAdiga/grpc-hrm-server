package db

import (
	"context"

	"github.com/google/uuid"
)

// Store is a base interface
// defines set of method to be implemented by different stores
// interface holding different database queries to be performed
type Store interface {
	CreateOrganization(ctx context.Context, param CreateOrganizationParam) (Organization, error)
	FindOrganizationByName(ctx context.Context, name string) (Organization, error)
	CreateUser(ctx context.Context, arg CreateUserParam) (uuid.UUID, error)
	FindUserByName(ctx context.Context, userName string) (User, error)
	CreateRole(ctx context.Context, arg CreateRoleParam) (uuid.UUID, error)
	FindRoleByOrganization(ctx context.Context, arg FindRoleByOrgParam) (Role, error)
	CreateEmployee(ctx context.Context, arg CreateEmployeeParam) (uuid.UUID, error)
	FindEmployeeByUnameAndOrg(ctx context.Context, arg FindEmployeeUnameAndOrgParam) (Employee, error)
}

// A compile time check to make sure Oueries implements Querier
// var _ Store = (*SQLStore)(nil)
