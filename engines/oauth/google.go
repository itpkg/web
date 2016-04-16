package oauth

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

/**
https://developers.google.com/identity/protocols/OAuth2WebServer
https://developers.google.com/identity/protocols/googlescopes
*/

//GoogleUser google user model
type GoogleUser struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Link    string `json:"link"`
	Picture string `json:"picture"`
}

//Google google credentials model
type Google struct {
	Web struct {
		ClientID     string   `json:"client_id"`
		ClientSecret string   `json:"client_secret"`
		RedirectURLS []string `json:"redirect_uris"`
	} `json:"web"`
}

//To to oauth2 credentials
func (p *Google) To() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     p.Web.ClientID,
		ClientSecret: p.Web.ClientSecret,
		RedirectURL:  p.Web.RedirectURLS[0],
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

//ReadGoogle read google credentials.
func ReadGoogle(f string) (*Google, error) {
	var g Google
	fd, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	dec := json.NewDecoder(fd)
	err = dec.Decode(&g)
	return &g, err
}
