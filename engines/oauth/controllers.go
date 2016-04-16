package oauth

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/itpkg/web/settings"
	"github.com/martini-contrib/render"
)

func getSignIn(r render.Render, sp settings.Provider) {
	var g Google
	if err := sp.Get("oauth.google", &g); err == nil {
		r.JSON(http.StatusOK, map[string]string{"google": g.To().AuthCodeURL("state")})
	} else {
		r.Text(http.StatusInternalServerError, err.Error())
	}
}

func postSignIn(r render.Render) {
	r.Text(http.StatusInternalServerError, "fuck")
}

//Mount mount to web
func (p *Engine) Mount(rt martini.Router) {
	rt.Get("/oauth/sign_in", getSignIn)
	rt.Post("/oauth/sign_in", postSignIn)
}
