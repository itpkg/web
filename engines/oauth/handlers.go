package oauth

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/itpkg/web/settings"
	"github.com/itpkg/web/token"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
)

//ADMIN admin role name
const ADMIN = "admin"

//GoogleHandler google credentials handler
func GoogleHandler(c martini.Context, sp settings.Provider, r render.Render) {
	var g Google
	if err := sp.Get("oauth.google", &g); err == nil {
		c.Map(&g)
	} else {
		r.Text(http.StatusInternalServerError, err.Error())
	}
}

//CurrentUserHandler current user handler
func CurrentUserHandler(req *http.Request, db *gorm.DB, c martini.Context, j *token.Jwt, r render.Render) {
	d, e := j.ParseFromRequest(req)
	var u User
	if e == nil {
		e = db.Where("uid = ?", d["id"]).First(&u).Error
	}
	if e == nil {
		if u.IsAvailable() {
			c.Map(&u)
		} else {
			r.Error(http.StatusUnauthorized)
		}

	} else {
		r.Text(http.StatusInternalServerError, e.Error())
	}
}

//AdminHandler admin handler
func AdminHandler(u *User, dao *Dao, r render.Render) {
	if !dao.Is(u.ID, ADMIN) {
		r.Error(http.StatusUnauthorized)
	}
}
