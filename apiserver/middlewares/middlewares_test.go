package middlewares_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/yanshiyason/noonde_platform/apiserver/models"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/httptest"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/suite"
)

type MiddlewareSuite struct {
	*suite.Action
}

func Test_MiddlewareSuite(t *testing.T) {
	envy.Load("./testdata/.env")
	action, err := suite.NewActionWithFixtures(buffalo.New(buffalo.Options{}), packr.NewBox("./testdata/fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	ms := &MiddlewareSuite{
		Action: action,
	}
	suite.Run(t, ms)
}

func loginCurrentUser(userEmail string, req *httptest.JSON) {
	claims := jwt.MapClaims{}
	user := models.User{}
	err := models.DB.Where("email = ?", userEmail).First(&user)
	claims["resource_id"] = user.ID
	claims["resource_type"] = "user"
	claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := envy.MustGet("JWT_SECRET")
	if err != nil {
		panic(err)
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	bearerToken := fmt.Sprintf("Bearer %s", tokenString)
	req.Headers["Authorization"] = bearerToken
}
