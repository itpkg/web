package base

import "github.com/go-martini/martini"

//Mount mount to web
func (p *Engine) Mount(rt martini.Router) {
	rt.Get("/locales/(:lang).json", getLocales)
}
