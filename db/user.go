package db

import (
	"context"

	"github.com/google/uuid"
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

const findUserByName = `
	SELECT * FROM users
	WHERE user_name = $1 LIMIT 1;
`

func (store *SQLStore) FindUserByName(ctx context.Context, userName string) (User, error) {
	row := store.db.QueryRowContext(ctx, findUserByName, userName)

	var u User

	err := row.Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.HashedPassword,
		&u.Address, &u.Email, &u.ContactNumber, &u.CreatedAt, &u.UpdatedAt,
	)

	return u, err
}

const findUserById = `
	SELECT * FROM users
	WHERE id = $1
`

func (store *SQLStore) FindUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := store.db.QueryRowContext(ctx, findUserById, id)

	var u User

	err := row.Scan(
		&u.ID, &u.FirstName, &u.LastName, &u.UserName, &u.HashedPassword,
		&u.Address, &u.Email, &u.ContactNumber, &u.CreatedAt, &u.UpdatedAt,
	)

	return u, err
}
