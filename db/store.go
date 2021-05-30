package db

import (
	"context"
)

// Store is a base interface
// defines set of method to be implemented by different stores
// interface holding different database queries to be performed
type Store interface {
	CreateOrganization(ctx context.Context, param CreateOrganizationParam) (Organization, error)
	FindOrganizationByName(ctx context.Context, name string) (Organization, error)
}

// A compile time check to make sure Oueries implements Querier
// var _ Store = (*db.SQLStore)(nil)
