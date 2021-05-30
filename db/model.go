package db

import (
	"time"

	"github.com/google/uuid"
)

/*
Model package is a database entity model used to store/retrive data from
database only
*/

// An organization table entity
type Organization struct {
	// organization id
	ID uuid.UUID
	// organization title
	Name string
	// Admin name
	CreatedBy string
	CreatorID uuid.UUID
	// state of organization
	Status uint16
	// Admin name
	UpdatedBy string
	UpdaterID uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
