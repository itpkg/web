package oauth

import (
	"fmt"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/itpkg/web/token"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
)

func getSignIn(g *Google, r render.Render) {
	r.JSON(http.StatusOK, map[string]string{"google": g.To().AuthCodeURL("state")})
}

func postSignIn(jwt *token.Jwt, dao *Dao, db *gorm.DB, g *Google, rdr render.Render, req *http.Request) {
	req.ParseForm()
	flag := req.Form.Get("type")
	code := req.Form.Get("code")
	var user *User
	var err error
	switch flag {
	case "google":
		var gu *GoogleUser
		gu, err = g.Parse(code)
		if err != nil {
			break
		}
		user, err = gu.Save(db)
	default:
		err = fmt.Errorf(fmt.Sprintf("Unsupported oauth %s", flag))
	}

	var tkn string
	if err == nil {
		tkn, err = jwt.New(
			map[string]interface{}{
				"id":      user.UID,
				"name":    user.Name,
				"isAdmin": dao.Is(user.ID, "admin"),
			}, 7*24*60)
	}

	if err == nil {
		rdr.Text(http.StatusOK, tkn)
	} else {
		rdr.Text(http.StatusInternalServerError, err.Error())
	}

}

//Mount mount to web
func (p *Engine) Mount(rt martini.Router) {
	rt.Get("/oauth/sign_in", GoogleHandler, getSignIn)
	rt.Post("/oauth/sign_in", GoogleHandler, postSignIn)
}
