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

const findPayrollByEmp = `
	SELECT * FROM payrolls 
	WHERE employee = $1
	LIMIT 1;
`

func (store *SQLStore) FindPayrollByEmp(ctx context.Context, id uuid.UUID) (Payroll, error) {
	row := store.db.QueryRowContext(ctx, findPayrollByEmp, id)

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
