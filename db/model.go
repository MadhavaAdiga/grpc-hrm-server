package db

import (
	"time"

	"github.com/google/uuid"
)

/*
Model package is a database entity model used to store/retrive data from
database only
*/

// users table entity
type User struct {
	ID             uuid.UUID
	FirstName      string
	LastName       string
	UserName       string
	HashedPassword string
	Address        string
	Email          string
	ContactNumber  uint32
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// organizations table entity
type Organization struct {
	ID        uuid.UUID // organization id
	Name      string    // organization title
	CreatorID uuid.UUID
	Status    uint16 // state of organization
	UpdaterID uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// role table entity
type Role struct {
	ID           uuid.UUID
	Name         string
	Active       bool
	Organization Organization // foreign key
	Permissions  []int32
	CreatedBy    uuid.UUID // admin id
	UpdatedBy    uuid.UUID // user id
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// employees table entity
type Employee struct {
	ID           uuid.UUID
	User         User
	Organization Organization
	Role         Role
	Status       int16
	CreateBy     uuid.UUID
	UpdatedBy    uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// payrolls table entity
type Payroll struct {
	ID        uuid.UUID
	Employee  Employee
	Ctc       int32
	Allowance int32
	CreateBy  uuid.UUID
	UpdatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
