package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Equipment struct {
	ID                 uuid.UUID `json:"id" db:"id"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
	User               uuid.UUID `json:"user" db:"user"`
	EquipedTo          uuid.UUID `json:"equiped_to" db:"equiped_to"`
	Name               string    `json:"name" db:"name"`
	Slot               string    `json:"slot" db:"slot"`
	PrimaryStat        string    `json:"primary_stat" db:"primary_stat"`
	PrimaryStatValue   int       `json:"primary_stat_value" db:"primary_stat_value"`
	SecondaryStat      string    `json:"secondary_stat" db:"secondary_stat"`
	SecondaryStatValue int       `json:"secondary_stat_value" db:"secondary_stat_value"`
	Level              int       `json:"level" db:"level"`
	Tier               string    `json:"tier" db:"tier"`
	Exp                int       `json:"exp" db:"exp"`
	ExpMax             int       `json:"exp_max" db:"exp_max"`
}

// String is not required by pop and may be deleted
func (e Equipment) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Equipment is not required by pop and may be deleted
type Equipment []Equipment

// String is not required by pop and may be deleted
func (e Equipment) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Equipment) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Name, Name: "Name"},
		&validators.StringIsPresent{Field: e.Slot, Name: "Slot"},
		&validators.StringIsPresent{Field: e.PrimaryStat, Name: "PrimaryStat"},
		&validators.IntIsPresent{Field: e.PrimaryStatValue, Name: "PrimaryStatValue"},
		&validators.StringIsPresent{Field: e.SecondaryStat, Name: "SecondaryStat"},
		&validators.IntIsPresent{Field: e.SecondaryStatValue, Name: "SecondaryStatValue"},
		&validators.IntIsPresent{Field: e.Level, Name: "Level"},
		&validators.StringIsPresent{Field: e.Tier, Name: "Tier"},
		&validators.IntIsPresent{Field: e.Exp, Name: "Exp"},
		&validators.IntIsPresent{Field: e.ExpMax, Name: "ExpMax"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Equipment) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Equipment) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
