package actions_test

import (
	"github.com/yanshiyason/noonde_platform/apiserver/models"
)

// NOTE: this is an unprotected endpoint.
// func (suite *ActionSuite) Test_UsersCreateHandler_Unauthorized() {
// 	res := suite.JSON("/users").Post(&models.User{})

// 	type E struct {
// 		Error string `json:"error"`
// 		Code  int    `json:"code"`
// 	}
// 	e := &E{}
// 	res.Bind(e)

// 	suite.Equal(401, res.Code)
// 	suite.Equal("token not found in request", e.Error)
// 	suite.Equal(401, e.Code)
// }

// NOTE: this is an unprotected endpoint.
// func (suite *ActionSuite) Test_UsersCreateHandler_Unauthorized_InvalidUserID() {
// 	req := suite.JSON("/users")
// 	loginCurrentUser("non-existant@example.org", req)

// 	res := req.Post(&models.User{})

// 	type E struct {
// 		Error string `json:"error"`
// 		Code  int    `json:"code"`
// 	}
// 	e := &E{}
// 	res.Bind(e)

// 	suite.Equal(401, res.Code)
// 	suite.Equal("invalid user ID", e.Error)
// 	suite.Equal(401, e.Code)
// }

func (suite *ActionSuite) Test_UsersCreateHandler() {
	// don't need to be logged in for this endpoint
	// suite.LoadFixture("current user")
	req := suite.JSON("/users")
	// currentUser := loginCurrentUser("user@example.co.jp", req)

	res := req.Post(&models.User{
		Name:                 "Mike Tyson",
		Password:             "mike",
		PasswordConfirmation: "mike",
		Email:                "mike@tyson.com",
	})
	suite.Equal(201, res.Code)

	response := &models.User{}
	res.Bind(response)
	suite.NotEmpty(response.ID)
	suite.NotEmpty(response.Name)
	suite.Empty(response.Password)
	suite.Empty(response.PasswordConfirmation)
}
