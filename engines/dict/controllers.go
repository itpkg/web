package dict

import (
	"net/http"

	"github.com/go-martini/martini"
	"github.com/itpkg/web/engines/oauth"
	"github.com/jinzhu/gorm"
	"github.com/martini-contrib/render"
)

func postIndex(dp Provider, req *http.Request, r render.Render) {
	req.ParseForm()

	if rs, err := dp.Query(req.FormValue("keyword")); err == nil {
		r.Text(http.StatusOK, rs)
	} else {
		r.Text(http.StatusInternalServerError, err.Error())
	}
}

func getIndex(dp Provider, r render.Render) {
	if ds, err := dp.List(); err == nil {
		r.JSON(http.StatusOK, ds)
	} else {
		r.Text(http.StatusInternalServerError, err.Error())
	}

}

func indexNotes(db *gorm.DB, u *oauth.User, r render.Render) {
	var notes []Note
	if err := db.Where("user_id = ?", u.ID).Find(&notes).Error; err == nil {
		r.JSON(http.StatusOK, notes)
	} else {
		r.Text(http.StatusInternalServerError, err.Error())
	}
}

func createNote(db *gorm.DB, u *oauth.User, r render.Render, req *http.Request) {
	req.ParseForm()
	n := Note{Title: req.FormValue("title"), Body: req.FormValue("body"), UserID: u.ID}
	if err := db.Create(&n).Error; err == nil {
		r.JSON(http.StatusOK, n)
	} else {
		r.Text(http.StatusInternalServerError, err.Error())
	}
}

func showNote(n *Note, r render.Render) {
	r.JSON(http.StatusOK, n)
}

func updateNote(db *gorm.DB, n *Note, r render.Render, req *http.Request) {
	req.ParseForm()
	n.Title = req.FormValue("title")
	n.Body = req.FormValue("body")
	if err := db.Save(n).Error; err == nil {
		r.JSON(http.StatusOK, n)
	} else {
		r.Text(http.StatusInternalServerError, err.Error())
	}
}

func destroyNote(db *gorm.DB, n *Note, r render.Render) {
	if err := db.Delete(n).Error; err == nil {
		r.JSON(http.StatusOK, n)
	} else {
		r.Text(http.StatusInternalServerError, err.Error())
	}
}

//Mount mount to web
func (p *Engine) Mount(rt martini.Router) {
	rt.Group("/dict", func(r martini.Router) {
		r.Get("", getIndex)
		r.Post("", postIndex)

		r.Get("/notes", indexNotes)
		r.Post("/notes", createNote)
		r.Get("/notes/:id", noteHandler, showNote)
		r.Post("/notes/:id", noteHandler, updateNote)
		r.Delete("/notes/:id", noteHandler, destroyNote)
	}, oauth.CurrentUserHandler)

}
