package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type parseFunc func(jwt.Keyfunc) (*jwt.Token, error)

//Provider token provider
type Provider interface {
	Set(kid string, key []byte, exp time.Duration) error
	Get(kid string) ([]byte, error)
	All() (map[string]int, error)
	Clear() error
}
