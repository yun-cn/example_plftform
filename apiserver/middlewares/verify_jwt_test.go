package middlewares_test

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	tokenauth "github.com/gobuffalo/mw-tokenauth"

	"github.com/yanshiyason/noonde_platform/apiserver/middlewares"
)

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func setupApp(suite *MiddlewareSuite) {
	app := buffalo.New(buffalo.Options{})
	suite.App = app

	app.Use(tokenauth.New(tokenauth.Options{
		SignMethod: jwt.SigningMethodHS256,
	}))
	app.Use(middlewares.VerifyJWT)

	app.GET("/", func(c buffalo.Context) error {
		return nil
	})
}

func (suite *MiddlewareSuite) Test_VerifyJWT_NoToken() {
	setupApp(suite)

	res := suite.JSON("/").Get()
	suite.Equal(401, res.Code)

	e := &ErrorResponse{}
	res.Bind(e)

	suite.Equal("token not found in request", e.Error)
	suite.Equal(401, e.Code)
}

func (suite *MiddlewareSuite) Test_VerifyJWT_InvalidAdmin() {
	setupApp(suite)
	req := suite.JSON("/")
	loginCurrentUser("not-found@example.co.jp", req)
	res := req.Get()

	suite.Equal(401, res.Code)

	e := &ErrorResponse{}
	res.Bind(e)

	suite.Equal("invalid user ID", e.Error)
	suite.Equal(401, e.Code)
}

func (suite *MiddlewareSuite) Test_VerifyJWT_ValidAdmin() {
	setupApp(suite)
	suite.LoadFixture("current user")

	req := suite.JSON("/")
	loginCurrentUser("user@example.co.jp", req)
	res := req.Get()

	suite.Equal(200, res.Code)
}
