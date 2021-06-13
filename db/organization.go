package db

import (
	"context"

	"github.com/google/uuid"
)

/*
	Organization DB service
	Provides abstraction for -
	CREATE,FIND
*/

const createOrganization = `
	INSERT INTO organizations (
		name,
		creator_id,
		status,
		updater_id
	) VALUES (
		$1,$2,$3,$4
	) RETURNING *
`

type CreateOrganizationParam struct {
	Name      string
	CreatorID uuid.UUID
	Status    uint16
}

func (store *SQLStore) CreateOrganization(ctx context.Context, arg CreateOrganizationParam) (Organization, error) {
	row := store.db.QueryRowContext(
		ctx, createOrganization, arg.Name, arg.CreatorID, arg.Status, arg.CreatorID,
	)

	var o Organization

	err := row.Scan(
		&o.ID, &o.Name, &o.CreatorID, &o.Status, &o.UpdaterID, &o.CreatedAt, &o.UpdatedAt,
	)

	return o, err
}

const findOrganization = `
	SELECT * FROM organizations
	WHERE name = $1 LIMIT 1
`

func (store *SQLStore) FindOrganizationByName(ctx context.Context, name string) (Organization, error) {
	row := store.db.QueryRowContext(ctx, findOrganization, name)

	var o Organization

	err := row.Scan(
		&o.ID, &o.Name, &o.CreatorID, &o.Status, &o.UpdaterID, &o.CreatedAt, &o.UpdatedAt,
	)

	return o, err
}
