package dict

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/itpkg/web/engines/oauth"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
)

func noteHandler(db *gorm.DB, dao *oauth.Dao, u *oauth.User, r render.Render, ps martini.Params, c martini.Context) {
	var nt Note
	if err := db.Where("id = ?", ps["id"]).First(&nt).Error; err == nil {
		if nt.UserID == u.ID || dao.Is(u.ID, oauth.ADMIN) {
			c.Map(&nt)
		} else {
			r.Text(http.StatusUnauthorized, err.Error())
		}
	} else {
		r.Text(http.StatusInternalServerError, err.Error())
	}
}
