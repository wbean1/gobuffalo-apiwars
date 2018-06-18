package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Exploration struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	User      uuid.UUID `json:"user" db:"user"`
	Zone      uuid.UUID `json:"zone" db:"zone"`
	Party     uuid.UUID `json:"party" db:"party"`
	FoundType string    `json:"found_type" db:"found_type"`
	FoundID   uuid.UUID `json:"found_id" db:"found_id"`
}

// String is not required by pop and may be deleted
func (e Exploration) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Explorations is not required by pop and may be deleted
type Explorations []Exploration

// String is not required by pop and may be deleted
func (e Explorations) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Exploration) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.FoundType, Name: "FoundType"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Exploration) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Exploration) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
