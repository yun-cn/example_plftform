package middlewares

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/buffalo"

	"github.com/yanshiyason/noonde_platform/apiserver/models"
)

// VerifyJWT make sure the admin exists in the DB and doesn't have too many tokens.
func VerifyJWT(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		// do some work before calling the next handler
		// Get Claims from JWT token from buffalo context
		claims := c.Value("claims").(jwt.MapClaims)
		UUID := claims["resource_id"].(string)
		Type := claims["resource_type"].(string)

		switch Type {
		case "user":
			if reason, err := handleUser(UUID); err != nil {
				c.LogField("error", err)
				return c.Error(401, errors.New(reason))
			}
		default:
			return c.Error(401, errors.New("Wrong jwt claims"))
		}

		return next(c)
	}
}

// handleUser returns true if valid, false and reason otherwise.
func handleUser(uuid string) (string, error) {
	count, err := models.DB.Where("id = ?", uuid).Count(&models.User{})
	if err != nil {
		// TODO log error
		return "Internal error", err
	}
	if count == 0 {
		// do some work after calling the next handler
		return "invalid user ID", errors.New("")
	}

	// TODO: do more work here
	// models.DB.Find(userID)

	return "", nil
}
