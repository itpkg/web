package base

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/itpkg/web/i18n"
	"github.com/martini-contrib/render"
	"golang.org/x/text/language"
)

func getLocales(lng *language.Tag, t *i18n.I18n, ps martini.Params, r render.Render) {
	r.JSON(http.StatusOK, t.Items(lng))
}
