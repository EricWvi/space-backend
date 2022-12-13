package service

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/space-backend/config"
	"net/http"
)

var (
	client = &http.Client{}
)

// Sign signs the identity with the specified secret.
func Sign(id uint) (tokenString string, err error) {
	// The usertoken content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
	})
	// Sign the usertoken with the specified secret.
	tokenString, err = token.SignedString([]byte(config.GetSecret()))
	return
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// ParseToken validates the usertoken with the specified secret,
// and returns the context if the usertoken was valid.
func ParseToken(tokenString string) (id uint, err error) {
	// Parse the usertoken.
	token, err := jwt.Parse(tokenString, secretFunc(config.GetSecret()))

	// Parse error.
	if err != nil {
		err = errors.New("token is invalid")
		return

		// Read the usertoken if it's valid.
	} else {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id = uint(claims["id"].(float64))
			// Other errors.
		} else {
			err = errors.New("token is invalid")
		}
	}
	return
}
