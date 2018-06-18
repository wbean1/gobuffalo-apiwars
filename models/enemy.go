package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Enemy struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	User      uuid.UUID `json:"user" db:"user"`
	Zone      uuid.UUID `json:"zone" db:"zone"`
	Name      string    `json:"name" db:"name"`
	Type      string    `json:"type" db:"type"`
	Level     int       `json:"level" db:"level"`
	Hp        int       `json:"hp" db:"hp"`
	Atk       int       `json:"atk" db:"atk"`
	Crit      int       `json:"crit" db:"crit"`
	Def       int       `json:"def" db:"def"`
	Armor     int       `json:"armor" db:"armor"`
}

// String is not required by pop and may be deleted
func (e Enemy) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Enemies is not required by pop and may be deleted
type Enemies []Enemy

// String is not required by pop and may be deleted
func (e Enemies) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Enemy) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Name, Name: "Name"},
		&validators.StringIsPresent{Field: e.Type, Name: "Type"},
		&validators.IntIsPresent{Field: e.Level, Name: "Level"},
		&validators.IntIsPresent{Field: e.Hp, Name: "Hp"},
		&validators.IntIsPresent{Field: e.Atk, Name: "Atk"},
		&validators.IntIsPresent{Field: e.Crit, Name: "Crit"},
		&validators.IntIsPresent{Field: e.Def, Name: "Def"},
		&validators.IntIsPresent{Field: e.Armor, Name: "Armor"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Enemy) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Enemy) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
