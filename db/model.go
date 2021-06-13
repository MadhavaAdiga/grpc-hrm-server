package db

import (
	"time"

	"github.com/google/uuid"
)

/*
Model package is a database entity model used to store/retrive data from
database only
*/

// User table entity
type User struct {
	ID             uuid.UUID
	FirstName      string
	LastName       string
	UserName       string
	HashedPassword string
	Address        string
	Email          string
	ContactNumber  uint32
	Employee_id    uuid.UUID
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// An organization table entity
type Organization struct {
	ID        uuid.UUID // organization id
	Name      string    // organization title
	CreatorID uuid.UUID
	Status    uint16 // state of organization
	UpdaterID uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role struct {
	ID           uuid.UUID
	Name         string
	Active       bool
	Organization uuid.UUID // foreign key
	Permissions  []int32
	CreatedBy    uuid.UUID // admin id
	UpdatedBy    uuid.UUID // user id
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
