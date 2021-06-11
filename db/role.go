package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

/*
	Role DB service
	Provides abstraction for -
	CREATE,FIND
*/

const createRole = `
	INSERT INTO roles (
		name         
		active       
		organization 
		permissions  
		createdBy    
	) VALUES (
		$1,$2,$3,$4,$5
	) RETURNING id
`

type CreateRoleParam struct {
	Name         string
	Active       bool
	Organization uuid.UUID
	Permissions  []int32
	CreatedBy    uuid.UUID
}

func (store *SQLStore) CreateRole(ctx context.Context, arg CreateRoleParam) (uuid.UUID, error) {
	row := store.db.QueryRowContext(ctx, createRole, arg.Name, arg.Active, arg.Organization, pq.Array(arg.Permissions), arg.CreatedBy)

	var id uuid.UUID

	err := row.Scan(&id)

	return id, err
}
