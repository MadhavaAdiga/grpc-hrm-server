package db

import (
	"context"

	"github.com/google/uuid"
)

const createEmployee = `
	INSERT INTO employees (
		user           
		organization 
		role         
		status     
		create_by    
	) VALUES(
		$1,$2,$3,$4,$5,$6
	) RETURNING id;
`

type CreateEmployeeParam struct {
	User_id         uuid.UUID
	Organization_id uuid.UUID
	Role_id         uuid.UUID
	Status          int16
	CreatedBy       uuid.UUID
}

func (store *SQLStore) CreateEmployee(ctx context.Context, arg CreateEmployeeParam) (uuid.UUID, error) {
	row := store.db.QueryRowContext(
		ctx, createEmployee, arg.User_id, arg.Organization_id, arg.Role_id, arg.Status, arg.CreatedBy,
	)

	var id uuid.UUID

	err := row.Scan(&id)

	return id, err
}

const findEmployeeUnameAndOrg = `
	EXPLAIN ANALYZE SELECT e.id, e."user",u.user_name, e.organization,o."name", e."role",r."name", e.status, e.create_by
	FROM employees e
	INNER JOIN users u  ON e."user" = u.id  
	LEFT JOIN roles r ON  e."role" = r.id
	JOIN organizations o ON o."name" =$1
	WHERE u.user_name = $2;
`

type FindEmployeeUnameAndOrgParam struct {
	OrganizationName string
	Username         string
}

func (store *SQLStore) FindEmployeeByUnameAndOrg(ctx context.Context, arg FindEmployeeUnameAndOrgParam) (Employee, error) {
	row := store.db.QueryRowContext(ctx, findEmployeeUnameAndOrg, arg.OrganizationName, arg.Username)

	var e Employee

	err := row.Scan(
		&e.Id,
		&e.User.ID,
		&e.User.UserName,
		&e.Organization.ID,
		&e.Organization.Name,
		&e.Role.ID,
		&e.Role.Name,
		&e.Status,
		&e.Create_by,
	)

	return e, err
}
