package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

/*
	Employee DB service
	Provides abstraction for -
		CREATE,FIND
*/

const createEmployee = `
	INSERT INTO employees (
		"user",           
		"organization", 
		"role",         
		status,     
		create_by,
		updated_by    
	) VALUES (
		$1,$2,$3,$4,$5,$6
	) RETURNING *;
`

type CreateEmployeeParam struct {
	UserId         uuid.UUID
	OrganizationId uuid.UUID
	RoleId         uuid.UUID
	Status         int16
	CreatedBy      uuid.UUID
}

func (store *SQLStore) CreateEmployee(ctx context.Context, arg CreateEmployeeParam) (Employee, error) {
	row := store.db.QueryRowContext(
		ctx, createEmployee, arg.UserId, arg.OrganizationId, arg.RoleId, arg.Status, arg.CreatedBy, arg.CreatedBy,
	)

	var e Employee

	err := row.Scan(
		&e.ID,
		&e.User.ID,
		&e.Organization.ID,
		&e.Role.ID,
		&e.Status,
		&e.CreateBy,
		&e.UpdatedBy,
		&e.CreatedAt,
		&e.UpdatedAt,
	)

	return e, err
}

const findEmployeeUnameAndOrg = `
	SELECT e.id, e."user",u.user_name, e.organization,o."name", e."role",r."name", e.status, e.create_by, e.created_at
	FROM
		employees e
		JOIN users u ON e."user" = u.id
		LEFT JOIN roles r ON e."role" = r.id
		JOIN organizations o ON o."name" = $1
	WHERE
		u.user_name = $2;
`

type FindEmployeeUnameAndOrgParam struct {
	OrganizationName string
	Username         string
}

func (store *SQLStore) FindEmployeeByUnameAndOrg(ctx context.Context, arg FindEmployeeUnameAndOrgParam) (Employee, error) {
	row := store.db.QueryRowContext(ctx, findEmployeeUnameAndOrg, arg.OrganizationName, arg.Username)

	var e Employee

	err := row.Scan(
		&e.ID,
		&e.User.ID,
		&e.User.UserName,
		&e.Organization.ID,
		&e.Organization.Name,
		&e.Role.ID,
		&e.Role.Name,
		&e.Status,
		&e.CreateBy,
		&e.CreatedAt,
	)

	return e, err
}

// find an employee on id joining organization name and role
const findAdminEmployee = `
	SELECT * FROM employees e
		JOIN organizations o ON o."name" = $1
		JOIN roles r ON r.id = e."role"
	WHERE
		e.id = $2;
`

type FindAdminEmployeeParam struct {
	OrganizationName string
	EmployeeId       uuid.UUID
}

// find an employee for the organization with roles
func (store SQLStore) FindAdminEmployee(ctx context.Context, arg FindAdminEmployeeParam) (Employee, error) {
	row := store.db.QueryRowContext(ctx, findAdminEmployee, arg.OrganizationName, arg.EmployeeId)

	var e Employee

	err := row.Scan(
		&e.ID,
		&e.User.ID,
		&e.Organization.ID,
		&e.Role.ID,
		&e.Status,
		&e.CreateBy,
		&e.UpdatedBy,
		&e.CreatedAt,
		&e.UpdatedAt,
		&e.Organization.ID, &e.Organization.Name, &e.Organization.CreatorID, &e.Organization.Status, &e.Organization.UpdaterID, &e.Organization.CreatedAt, &e.Organization.UpdatedAt,
		&e.Role.ID, &e.Role.Name, &e.Role.Active, &e.Role.Organization.ID, pq.Array(&e.Role.Permissions), &e.Role.CreatedBy, &e.Role.UpdatedBy, &e.Role.CreatedAt, &e.Role.UpdatedAt,
	)

	return e, err
}
