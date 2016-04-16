package oauth

import (
	"encoding/json"
	"os"

	"github.com/itpkg/web"
	"github.com/itpkg/web/engines/base"
	"github.com/jinzhu/gorm"

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

//Save save user to database
func (p *GoogleUser) Save(db *gorm.DB) (*base.User, error) {

	var u base.User
	err := db.Where("provider_id = ? AND provider_type = ?", "google", p.ID).First(&u).Error
	u.Email = p.Email
	u.Name = p.Name
	u.Logo = p.Picture
	u.Home = p.Link

	if err == nil {
		db.Save(&u)
	} else {
		u.UID = web.UUID()
		u.ProviderID = p.ID
		u.ProviderType = "google"
		db.Create(&u)
	}

	return &u, nil
}

//Google google credentials model
type Google struct {
	Web struct {
		ClientID     string   `json:"client_id"`
		ClientSecret string   `json:"client_secret"`
		RedirectURLS []string `json:"redirect_uris"`
	} `json:"web"`
}

//Parse parse user from request
func (p *Google) Parse(code string) (*GoogleUser, error) {
	cfg := p.To()

	tok, err := cfg.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}
	cli := cfg.Client(oauth2.NoContext, tok)
	res, err := cli.Get("https://www.googleapis.com/oauth2/v1/userinfo?alt=json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var gu GoogleUser
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&gu)
	return &gu, err
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
