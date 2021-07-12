package db

import (
	"context"

	"github.com/google/uuid"
)

/*
	Payroll DB service
	Provides abstraction for -
		CREATE,FIND
*/

const createPayroll = `
	INSERT INTO payrolls (
		"employee",
		ctc,
		allowance,
		create_by,
		updated_by
	) VALUES (
		$1,$2,$3,$4,$5
	) RETURNING *;
`

type CreatePayrollParam struct {
	Employee  uuid.UUID
	Ctc       int32
	Allowance int32
	CreatedBy uuid.UUID
}

func (store *SQLStore) CreatePayroll(ctx context.Context, arg CreatePayrollParam) (Payroll, error) {
	row := store.db.QueryRowContext(
		ctx, createPayroll, arg.Employee, arg.Ctc, arg.Allowance, arg.CreatedBy, arg.CreatedBy,
	)

	var p Payroll

	err := row.Scan(
		&p.ID,
		&p.Employee.ID,
		&p.Ctc,
		&p.Allowance,
		&p.CreateBy,
		&p.UpdatedBy,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}

const findPayrollByEmpId = `
	SELECT * FROM payrolls p
	JOIN employees e ON e.id = p.employee
		WHERE employee = $1;
`

func (store *SQLStore) FindPayrollByEmpID(ctx context.Context, id uuid.UUID) (Payroll, error) {
	row := store.db.QueryRowContext(ctx, findPayrollByEmpId, id)

	var p Payroll

	err := row.Scan(
		&p.ID,
		&p.Employee.ID,
		&p.Ctc,
		&p.Allowance,
		&p.CreateBy,
		&p.UpdatedBy,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Employee.ID, &p.Employee.User.ID, &p.Employee.Organization.ID, &p.Employee.Role.ID, &p.Employee.Status, &p.Employee.CreateBy, &p.Employee.UpdatedBy, &p.Employee.CreatedAt, &p.Employee.UpdatedAt,
	)

	return p, err
}

const findPayrollByEmpName = `
	SELECT * FROM payrolls p
	JOIN employees e ON e.id = p.employee
		WHERE e."user" = (SELECT id from users where user_name = $1);
`

func (store *SQLStore) FindPayrollByEmpName(ctx context.Context, name string) (Payroll, error) {
	row := store.db.QueryRowContext(ctx, findPayrollByEmpName, name)

	var p Payroll

	err := row.Scan(
		&p.ID,
		&p.Employee.ID,
		&p.Ctc,
		&p.Allowance,
		&p.CreateBy,
		&p.UpdatedBy,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.Employee.ID, &p.Employee.User.ID, &p.Employee.Organization.ID, &p.Employee.Role.ID, &p.Employee.Status, &p.Employee.CreateBy, &p.Employee.UpdatedBy, &p.Employee.CreatedAt, &p.Employee.UpdatedAt,
	)

	return p, err
}
