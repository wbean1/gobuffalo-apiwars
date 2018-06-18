package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Character struct {
	ID            uuid.UUID `json:"id" db:"id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
	User          User      `has_one:"user" fk_id:"id"`
	Name          string    `json:"name" db:"name"`
	Class         string    `json:"class" db:"class"`
	Level         int       `json:"level" db:"level"`
	Tier          string    `json:"tier" db:"tier"`
	Experience    int       `json:"experience" db:"experience"`
	ExperienceMax int       `json:"experience_max" db:"experience_max"`
	AtkBase       int       `json:"atk_base" db:"atk_base"`
	CritBase      int       `json:"crit_base" db:"crit_base"`
	DefBase       int       `json:"def_base" db:"def_base"`
	ArmorBase     int       `json:"armor_base" db:"armor_base"`
	HpBase        int       `json:"hp_base" db:"hp_base"`
	Party         Party     `belongs_to:"party"`
        PartyID       uuid.UUID `json:"party_id" db:"party_id"`
}

// String is not required by pop and may be deleted
func (c Character) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Characters is not required by pop and may be deleted
type Characters []Character

// String is not required by pop and may be deleted
func (c Characters) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Character) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name"},
		&validators.StringIsPresent{Field: c.Class, Name: "Class"},
		&validators.IntIsPresent{Field: c.Level, Name: "Level"},
		&validators.StringIsPresent{Field: c.Tier, Name: "Tier"},
		&validators.IntIsPresent{Field: c.Experience, Name: "Experience"},
		&validators.IntIsPresent{Field: c.ExperienceMax, Name: "ExperienceMax"},
		&validators.IntIsPresent{Field: c.AtkBase, Name: "AtkBase"},
		&validators.IntIsPresent{Field: c.CritBase, Name: "CritBase"},
		&validators.IntIsPresent{Field: c.DefBase, Name: "DefBase"},
		&validators.IntIsPresent{Field: c.ArmorBase, Name: "ArmorBase"},
		&validators.IntIsPresent{Field: c.HpBase, Name: "HpBase"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Character) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Character) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
