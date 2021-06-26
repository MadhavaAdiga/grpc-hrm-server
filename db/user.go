package db

import (
	"context"
)

/*
	User DB service
	Provides abstraction for -
		CREATE,FIND
*/

const createUser = `
	INSERT INTO users (
		first_name,
		last_name,
		user_name,
		hashed_password,
		address,
		email,
		contact_number
	) VALUES (
		$1,$2,$3,$4,$5,$6,$7
	) RETURNING *;
`

type CreateUserParam struct {
	FirstName      string
	LastName       string
	UserName       string
	HashedPassword string
	Address        string
	Email          string
	ContactNumber  uint32
}

func (store *SQLStore) CreateUser(ctx context.Context, arg CreateUserParam) (User, error) {
	row := store.db.QueryRowContext(ctx, createUser, arg.FirstName, arg.LastName, arg.UserName,
		arg.HashedPassword, arg.Address, arg.Email, arg.ContactNumber)

	var u User

	err := row.Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.HashedPassword,
		&u.Address, &u.Email, &u.ContactNumber, &u.CreatedAt, &u.UpdatedAt,
	)

	return u, err
}

const findUser = `
	SELECT * FROM users
	WHERE user_name = $1 LIMIT 1;
`

func (store *SQLStore) FindUserByName(ctx context.Context, userName string) (User, error) {
	row := store.db.QueryRowContext(ctx, findUser, userName)

	var u User

	err := row.Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.HashedPassword,
		&u.Address, &u.Email, &u.ContactNumber, &u.CreatedAt, &u.UpdatedAt,
	)

	return u, err
}
