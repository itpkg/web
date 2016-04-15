package base

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/itpkg/web/i18n"
	"github.com/itpkg/web/settings"
	"github.com/martini-contrib/render"
	"golang.org/x/text/language"
)

//Site site info model
type Site struct {
	Lang        string   `json:"lang"`
	Title       string   `json:"title"`
	SubTitle    string   `json:"subTitle"`
	Description string   `json:"description"`
	Keywords    string   `json:"keywords"`
	Copyright   string   `json:"copyright"`
	Links       []string `json:"links"`
}

func getSiteInfo(lng *language.Tag, sp settings.Provider, r render.Render, lg *log.Logger) {
	var si Site
	if err := sp.Get(fmt.Sprintf("%s://site.info", lng.String()), &si); err != nil {
		lg.Print(err)
		si.Lang = lng.String()
		si.Links = make([]string, 0)
	}
	r.JSON(http.StatusOK, si)
}

func getLocales(t *i18n.I18n, ps martini.Params, r render.Render) {
	lng := i18n.Match(ps["lang"])
	r.JSON(http.StatusOK, t.Items(&lng))
}
