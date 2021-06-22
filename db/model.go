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
	Employee_id    uuid.UUID
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
	Id           uuid.UUID
	User         User
	Organization Organization
	Role         Role
	Status       int16
	Create_by    uuid.UUID
	Updated_by   uuid.UUID
	Created_at   time.Time
	Updated_at   time.Time
}

// payrolls table entity
type Payroll struct {
	Id         uuid.UUID
	Employee   Employee
	Ctc        int32
	Allowance  int32
	Create_by  uuid.UUID
	Updated_by uuid.UUID
	Created_at time.Time
	Updated_at time.Time
}
