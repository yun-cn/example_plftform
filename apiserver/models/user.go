package models

import (
	"encoding/json"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/pkg/errors"
)

// User is a user.
type User struct {
	ID                   uuid.UUID  `json:"id"                            db:"id"`
	CreatedAt            time.Time  `json:"created_at"                    db:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at"                    db:"updated_at"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty"          db:"deleted_at"`
	Name                 string     `json:"name"                          db:"name"`
	Email                string     `json:"email"                         db:"email"`
	PasswordHash         string     `json:"-"                             db:"password_hash"`
	Password             string     `json:"password,omitempty"            db:"-"`
	PasswordConfirmation string     `json:"password_confirmation,omitempty" db:"-"`
}

// String Stringer
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users User collection
type Users []User

// String Stringer
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Create create
func (u *User) Create(tx *pop.Connection) (*validate.Errors, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return validate.NewErrors(), errors.WithStack(err)
	}

	u.PasswordHash = string(hash)
	u.Email = strings.ToLower(u.Email)

	verrs, err := tx.ValidateAndCreate(u)
	if err != nil {
		return verrs, err
	}

	return verrs, nil
}

func (u *User) BeforeCreate(tx *pop.Connection) error {
	var err error
	u.ID, err = uuid.NewV1()
	return err
}

func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.EmailIsPresent{Field: u.Email, Name: "Email"},
		&validators.StringIsPresent{Field: u.Name, Name: "Name"},
		&validators.StringIsPresent{Field: u.PasswordHash, Name: "PasswordHash"},
		&validators.FuncValidator{
			Field:   u.Email,
			Name:    "Email",
			Message: "%s is already taken.",
			Fn: func() bool {
				var taken bool
				var err error
				q := tx.Where("email = ?", u.Email)
				if u.ID != uuid.Nil {
					// it's only taken if it belongs to another record
					q = q.Where("id != ?", u.ID)
				}
				taken, err = q.Exists(u)
				if err != nil {
					return false
				}
				return !taken
			},
		},
	), nil
}

func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	// var err error
	// verrs, err := u.Validate(tx)
	// if err != nil {
	// 	return verrs, errors.WithStack(err)
	// }
	// if verrs.HasAny() {
	// 	return verrs, nil
	// }
	// only during create:
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Password, Name: "Password"},
		&validators.StringsMatch{Field: u.Password, Field2: u.PasswordConfirmation, Name: "Password", Message: "Password does not match confirmation."},
	), nil
}

func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	verrs, err := u.Validate(tx)
	if err != nil {
		return verrs, errors.WithStack(err)
	}
	// if verrs.HasAny() {

	// }

	// err = tx.Update(u)
	// return nil, err
	return verrs, nil
}
