package i18n

import (
	"net/http"

	"github.com/go-martini/martini"
	"golang.org/x/text/language"
)

//Languages list
var Languages = []language.Tag{language.SimplifiedChinese, language.AmericanEnglish}

var matcher = language.NewMatcher(Languages)

//Match the default language
func Match(lang string) language.Tag {
	lng, _, _ := matcher.Match(language.Make(lang))
	return lng
}

//LangHandler parse locale from http request
func LangHandler(c martini.Context, req *http.Request, res http.ResponseWriter) {
	const key = "locale"
	// 1. Check URL arguments.
	lang := req.URL.Query().Get(key)
	// 2. Get language information from cookies.
	if len(lang) == 0 {
		if ck, err := req.Cookie(key); err == nil {
			lang = ck.Value
		}
	}
	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := req.Header.Get("Accept-Language")
		if len(al) > 4 {
			lang = al[:5] // Only compare first 5 letters.
		}
	}

	lng := Match(lang)

	// if !strings.EqualFold(lang, lng.String()) {
	// 	write = true
	// }
	// if write {
	//http.SetCookie(res, &http.Cookie{Name: key, Value: lng.String(), Path: "/", MaxAge: 1<<31 - 1})
	// }

	c.Map(&lng)
}
