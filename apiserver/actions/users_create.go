package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/yanshiyason/noonde_platform/apiserver/models"
)

// UsersCreateHandler is a handler to create an user
func UsersCreateHandler(c buffalo.Context) error {
	user := &models.User{}

	// Bind attributes to the incoming json
	if err := c.Bind(user); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	// Func to evaluate if the creation succeeds
	// AfterCreate(*Connection) error
	// user.AfterCreateFn = func(tx *pop.Connection) {
	// 	Worker.UploadUserAvatarJob(user.ID, user.Avatar)
	// 	Worker.IndexUserJob(user.ID)
	// }
	// Validate the data from the json
	verrs, err := user.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		return c.Render(400, r.JSON(verrs))
	}

	newUser := models.User{}
	tx.Where("email = ?", user.Email).Last(&newUser)

	return c.Render(201, r.JSON(&newUser))
}
