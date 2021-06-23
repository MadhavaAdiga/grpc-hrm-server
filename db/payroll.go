package db

import (
	"context"

	"github.com/google/uuid"
)

const createPayroll = `
	INSERT INTO payrolls (
		employee
		ctc
		allowance
		create_by
	) VALUES (
		$1,$2,$3,$4
	) RETURNING id;
`

type CreatePayrollParam struct {
	Employee
	Ctc        int32
	Allowance  int32
	Created_by uuid.UUID
}

func (store *SQLStore) CreatePayroll(ctx context.Context, arg CreatePayrollParam) (uuid.UUID, error) {
	row := store.db.QueryRowContext(ctx, createPayroll, arg.Employee, arg.Ctc, arg.Allowance, arg.Created_by)

	var id uuid.UUID

	err := row.Scan(&id)

	return id, err
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
		&p.Id,
		&p.Employee.Id,
		&p.Ctc,
		&p.Allowance,
		&p.Create_by,
		&p.Updated_by,
		&p.Created_at,
		&p.Updated_at,
	)

	return p, err
}
