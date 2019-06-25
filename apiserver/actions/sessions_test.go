package actions_test

import (
	"github.com/yanshiyason/noonde_platform/apiserver/actions"
)

func (suite *ActionSuite) Test_SessionsCreateHandler_CorrectPassword() {
	// don't need to be logged in for this endpoint
	suite.LoadFixture("mike tyson")
	req := suite.JSON("/sign_in")
	// // currentUser := loginCurrentUser("user@example.co.jp", req)

	res := req.Post(&actions.SignInParams{
		Email:    "mike@tyson.com",
		Password: "password",
	})

	response := &actions.SignInParams{}
	res.Bind(response)

	suite.Equal(200, res.Code)
	suite.Equal("True!", response.Email)

	// response := &models.User{}
	// res.Bind(response)
	// suite.NotEmpty(response.ID)
	// suite.NotEmpty(response.Name)
	// suite.Empty(response.Password)
	// suite.Empty(response.PasswordConfirmation)
}

func (suite *ActionSuite) Test_SessionsCreateHandler_WrongPassword() {
	// don't need to be logged in for this endpoint
	suite.LoadFixture("mike tyson")
	req := suite.JSON("/sign_in")
	// // currentUser := loginCurrentUser("user@example.co.jp", req)

	res := req.Post(&actions.SignInParams{
		Email:    "mike@tyson.com",
		Password: "wrong password",
	})

	response := &actions.SignInParams{}
	res.Bind(response)

	suite.Equal(400, res.Code)
	// suite.Equal("True!", response.Email)

	// response := &models.User{}
	// res.Bind(response)
	// suite.NotEmpty(response.ID)
	// suite.NotEmpty(response.Name)
	// suite.Empty(response.Password)
	// suite.Empty(response.PasswordConfirmation)
}
