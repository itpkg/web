package oauth

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/itpkg/web/settings"
	"github.com/martini-contrib/render"
)

//GoogleHandler google credentials handler
func GoogleHandler(c martini.Context, sp settings.Provider, r render.Render) {
	var g Google
	if err := sp.Get("oauth.google", &g); err == nil {
		c.Map(&g)
	} else {
		r.Text(http.StatusInternalServerError, err.Error())
	}
}
