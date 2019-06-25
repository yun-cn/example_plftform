package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
	"github.com/yanshiyason/noonde_platform/apiserver/models"
	"golang.org/x/crypto/bcrypt"
)

// SignInParams ..
type SignInParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SessionsCreateHandler is the handler to sign a user in
func SessionsCreateHandler(c buffalo.Context) error {
	params := &SignInParams{}

	// Bind attributes to the incoming json
	if err := c.Bind(params); err != nil {
		return errors.WithStack(err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	c.LogField("HASH", string(hash))
	c.LogField("PWD", params.Password)

	user := models.User{}
	err = models.DB.Where("email = ?", params.Email).First(&user)
	if err != nil {
		panic(err)
	}

	ok := comparePasswords(user.PasswordHash, params.Password)

	if ok {
		return c.Render(200, r.JSON(&SignInParams{Email: "True!"}))
	}

	return c.Render(400, r.JSON(&SignInParams{Email: "Wrong!"}))
}

// SessionsDeleteHandler is the handler to sign a user in
func SessionsDeleteHandler(c buffalo.Context) error {
	return c.Render(200, r.JSON(&SignInParams{}))
}

func comparePasswords(hashedPwd string, plainPwd string) bool {
	hashed := []byte(hashedPwd)
	plain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(hashed, plain)

	return err == nil
}
