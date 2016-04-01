package base

import "github.com/go-martini/martini"

//Mount mount to web
func (p *Engine) Mount(rt martini.Router) {
	rt.Get("/", getHome)
	rt.Get("/users/sign-up", getUsersSignUp)
	rt.Get("/users/confirm", getUsersConfirm)
	rt.Get("/users/sign-in", getUsersSignIn)
	rt.Get("/users/unlock", getUsersUnlock)
}
