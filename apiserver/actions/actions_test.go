package actions_test

import (
	"fmt"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/httptest"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/suite"

	"github.com/yanshiyason/noonde_platform/apiserver/actions"
	"github.com/yanshiyason/noonde_platform/apiserver/models"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(actions.App(), packr.New("Test_ActionSuite", "./testdata/fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}

func loginCurrentUser(userEmail string, req *httptest.JSON) *models.User {
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

	return &user
}
