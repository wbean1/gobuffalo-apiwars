package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Zone struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	User          uuid.UUID `json:"user" db:"user"`
	Name          string    `json:"name" db:"name"`
	Type          string    `json:"type" db:"type"`
	Level         int       `json:"level" db:"level"`
	Tier          string    `json:"tier" db:"tier"`
	ExploresCount int       `json:"explores_count" db:"explores_count"`
	ExploresMax   int       `json:"explores_max" db:"explores_max"`
}

// String is not required by pop and may be deleted
func (z Zone) String() string {
	jz, _ := json.Marshal(z)
	return string(jz)
}

// Zones is not required by pop and may be deleted
type Zones []Zone

// String is not required by pop and may be deleted
func (z Zones) String() string {
	jz, _ := json.Marshal(z)
	return string(jz)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (z *Zone) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: z.Name, Name: "Name"},
		&validators.StringIsPresent{Field: z.Type, Name: "Type"},
		&validators.IntIsPresent{Field: z.Level, Name: "Level"},
		&validators.StringIsPresent{Field: z.Tier, Name: "Tier"},
		&validators.IntIsPresent{Field: z.ExploresCount, Name: "ExploresCount"},
		&validators.IntIsPresent{Field: z.ExploresMax, Name: "ExploresMax"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (z *Zone) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (z *Zone) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
