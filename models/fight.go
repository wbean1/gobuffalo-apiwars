package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Fight struct {
	ID                   uuid.UUID `json:"id" db:"id"`
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time `json:"updated_at" db:"updated_at"`
	User                 uuid.UUID `json:"user" db:"user"`
	Enemy                uuid.UUID `json:"enemy" db:"enemy"`
	Party                uuid.UUID `json:"party" db:"party"`
	Result               string    `json:"result" db:"result"`
	FoundType            string    `json:"found_type" db:"found_type"`
	FoundID              uuid.UUID `json:"found_id" db:"found_id"`
	ExperienceAdjustment int       `json:"experience_adjustment" db:"experience_adjustment"`
}

// String is not required by pop and may be deleted
func (f Fight) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Fights is not required by pop and may be deleted
type Fights []Fight

// String is not required by pop and may be deleted
func (f Fights) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *Fight) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: f.Result, Name: "Result"},
		&validators.StringIsPresent{Field: f.FoundType, Name: "FoundType"},
		&validators.IntIsPresent{Field: f.ExperienceAdjustment, Name: "ExperienceAdjustment"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *Fight) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *Fight) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
