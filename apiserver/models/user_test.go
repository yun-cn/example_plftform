package models_test

import (
	"github.com/yanshiyason/noonde_platform/apiserver/models"
)

type User = models.User

func (ms *ModelSuite) Test_User_Create() {
	var err error
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	user := &User{Email: "user@example.com", Name: "Rob Pike", Password: "password", PasswordConfirmation: "password"}

	verrs, err := user.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny(), verrs)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
	var err error
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	user := &User{}
	verrs, err := user.Create(ms.DB)
	ms.NoError(err)
	ms.NotNil(verrs)
	ms.True(verrs.HasAny())

	ms.Contains(verrs.Get("email"), "Email does not match the email format.")
	ms.Contains(verrs.Get("name"), "Name can not be blank.")

	user = &User{Email: "invalid@email"}

	verrs, err = user.Create(ms.DB)
	ms.NoError(err)
	ms.Contains(verrs.Get("email"), "Email does not match the email format.")

	// Make sure no users were persisted to the db
	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationDuplicateEmail() {
	var err error
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	params1 := &User{Email: "user@example.com", Name: "Rob Pike", Password: "password", PasswordConfirmation: "password"}
	params2 := &User{Email: "user@example.com", Name: "Rob Pike", Password: "password", PasswordConfirmation: "password"}

	verrs, err := params1.Create(ms.DB)
	ms.NoError(err)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	// fails with duplicate email
	verrs, err = params2.Create(ms.DB)
	ms.NoError(err)
	ms.NotNil(verrs)
	ms.Contains(verrs.Get("email"), "user@example.com is already taken.")

	// does not create a record
	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_PasswordConfirmation() {
	var err error
	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	user := &User{Email: "user@example.com", Name: "Rob Pike", Password: "wrong", PasswordConfirmation: "password"}
	errs, err := user.Create(ms.DB)
	ms.NoError(err)
	ms.Contains(errs.Get("password"), "Password does not match confirmation.")

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}
