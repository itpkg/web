package token

import "github.com/dgrijalva/jwt-go"

type parseFunc func(jwt.Keyfunc) (*jwt.Token, error)

//Provider token provider
type Provider interface {
	Set(kid string, key []byte) error
	Get(kid string) ([]byte, error)
}
