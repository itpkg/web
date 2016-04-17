package token

import (
	"crypto/rand"
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pborman/uuid"
)

//Jwt jwt helper
type Jwt struct {
	Provider Provider
	Key      []byte
}

//ParseFromRequest parse token from request
func (p *Jwt) ParseFromRequest(req *http.Request) (map[string]interface{}, error) {
	return p.parse(func(fn jwt.Keyfunc) (*jwt.Token, error) {
		return jwt.ParseFromRequest(req, fn)
	})
}

//Parse parse token from string
func (p *Jwt) Parse(str string) (map[string]interface{}, error) {
	return p.parse(func(fn jwt.Keyfunc) (*jwt.Token, error) {
		return jwt.Parse(str, fn)
	})
}

func (p *Jwt) parse(fn parseFunc) (map[string]interface{}, error) {
	token, err := fn(func(token *jwt.Token) (interface{}, error) {
		key, err := p.Provider.Get(token.Header["kid"].(string))
		if err == nil {
			return append(p.Key, key...), nil
		}
		return nil, nil
	})
	if err == nil {
		if token.Valid {
			//delete(token.Claims, "exp")
			return token.Claims, nil
		}
		return nil, errors.New("token is not valid")

	}
	return nil, err
}

//New generate token
func (p *Jwt) New(data map[string]interface{}, exp time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS512)
	for k, v := range data {
		token.Claims[k] = v
	}
	token.Claims["exp"] = time.Now().Add(exp).Unix()
	kid := uuid.New()
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	if err := p.Provider.Set(kid, key, exp); err != nil {
		return "", err
	}

	token.Header["kid"] = kid
	return token.SignedString(append(p.Key, key...))
}
