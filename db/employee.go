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

/*
SELECT e.id, e."user", e.organization, e."role", e.status, e.payroll, e.create_by, e.updated_by, e.created_at, e.updated_at
	FROM employees e
	LEFT JOIN users u  ON e."user" = u.id
	LEFT JOIN roles r ON  e."role" = r.id
	WHERE u.user_name = $1 AND  r."name" = $1;
*/
// const findEmployeeByUserNameAndRole = `
// 	SELECT e.id, e."user", e.organization, e."role", e.status, e.payroll, e.create_by, e.updated_by, e.created_at, e.updated_at
// 	FROM employees e
// 	LEFT JOIN users u  ON e."user" = u.id
// 	LEFT JOIN roles r ON  e."role" = r.id
// 	WHERE u.user_name = $1
// 	AND  r."name" = $2;
// `

// type FindEmployeeByUnameAndRole struct {
// 	Username string
// 	role     string
// }
