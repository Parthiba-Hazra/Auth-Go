package models

import (
	"github.com/google/uuid"
)

// user model (The password's json value is "-" because we dont want to returned the password in any json responce)
type User struct {
	UID      uuid.UUID `db:""uid" json:"uid,omitempty"`
	Name     string    `db:"name" json:"email" validate:"required"`
	Email    string    `db:"email" json:"email" validate:"required"`
	Password string    `db:"password" json:"-" validate:"required"`
	ImageURL string    `db:"image_url" json:"imageurl"`
	Website  string    `db:"website" json:"website"`
}
