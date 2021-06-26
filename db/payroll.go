package db

import (
	"context"

	"github.com/google/uuid"
)

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

const findPayroll = `
	SELECT * FROM payrolls 
	WHERE employee = $1
	LIMIT 1;
`

func (store *SQLStore) FindPayroll(ctx context.Context, id uuid.UUID) (Payroll, error) {
	row := store.db.QueryRowContext(ctx, findPayroll, id)

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
