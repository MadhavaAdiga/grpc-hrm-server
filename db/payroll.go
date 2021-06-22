package db

import "github.com/google/uuid"

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

type CreatePayroll struct {
	employee  uuid.UUID
	ctc       int32
	allowance int32
	create_by uuid.UUID
}
