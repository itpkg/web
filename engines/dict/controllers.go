package dict

import (
	"net/http"

	"github.com/go-martini/martini"
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

func indexNotes() {

}

func createNote() {

}

func showNote() {

}

func updateNote() {

}
func destroyNote() {

}

//Mount mount to web
func (p *Engine) Mount(rt martini.Router) {
	rt.Get("/dict", getIndex)
	rt.Post("/dict", postIndex)

	rt.Get("/dict/notes", indexNotes)
	rt.Post("/dict/notes", createNote)
	rt.Get("/dict/notes/:id", showNote)
	rt.Post("/dict/notes/:id", updateNote)
	rt.Delete("/dict/notes/:id", destroyNote)
}
