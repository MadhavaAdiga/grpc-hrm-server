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
		"name",        
		active,
		organization,
		"permissions ",
		created_by,
		updated_by    
	) VALUES (
		$1,$2,$3,$4,$5,$6
	) RETURNING *;
`

type CreateRoleParam struct {
	Name         string
	Active       bool
	Organization uuid.UUID
	Permissions  []int32
	CreatedBy    uuid.UUID
}

func (store *SQLStore) CreateRole(ctx context.Context, arg CreateRoleParam) (Role, error) {
	row := store.db.QueryRowContext(
		ctx, createRole, arg.Name, arg.Active, arg.Organization, pq.Array(arg.Permissions), arg.CreatedBy, arg.CreatedBy,
	)

	var r Role

	err := row.Scan(
		&r.ID, &r.Name, &r.Active, &r.Organization.ID, pq.Array(&r.Permissions),
		&r.CreatedBy, &r.UpdatedBy, &r.CreatedAt, &r.UpdatedAt,
	)

	return r, err
}

const findRoleByOrganizationID = `
	SELECT * FROM roles 
	WHERE "name" =$1 AND "organization" =$2 
	LIMIT 1;
`

type FindRoleByOrgIDParam struct {
	Name         string
	Organization uuid.UUID
}

func (store *SQLStore) FindRoleByOrganizationID(ctx context.Context, arg FindRoleByOrgIDParam) (Role, error) {
	row := store.db.QueryRowContext(ctx, findRoleByOrganizationID, arg.Name, arg.Organization)

	var r Role

	err := row.Scan(
		&r.ID, &r.Name, &r.Active, &r.Organization.ID, pq.Array(&r.Permissions),
		&r.CreatedBy, &r.UpdatedBy, &r.CreatedAt, &r.UpdatedAt,
	)

	return r, err
}

const findRoleByOrganizationName = `
	SELECT * FROM roles 
	WHERE "name" =$1 AND "organization" = (SELECT id FROM organizations WHERE "name" = $2) 
	LIMIT 1;
`

type FindRoleByOrgNameParam struct {
	Name             string
	OrganizationName string
}

func (store *SQLStore) FindRoleByOrganizationName(ctx context.Context, arg FindRoleByOrgNameParam) (Role, error) {
	row := store.db.QueryRowContext(ctx, findRoleByOrganizationName, arg.Name, arg.OrganizationName)

	var r Role

	err := row.Scan(
		&r.ID, &r.Name, &r.Active, &r.Organization.ID, pq.Array(&r.Permissions),
		&r.CreatedBy, &r.UpdatedBy, &r.CreatedAt, &r.UpdatedAt,
	)

	return r, err
}
