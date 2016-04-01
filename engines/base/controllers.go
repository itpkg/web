package base

import (
	"net/http"

	"github.com/martini-contrib/render"
)

func getHome(r render.Render) {
	r.HTML(http.StatusOK, "home", nil)
}

func getUsersSignUp(r render.Render) {
	r.HTML(http.StatusOK, "users/sign-up", nil)
}

func getUsersConfirm(r render.Render) {
	r.HTML(http.StatusOK, "users/confirm", nil)
}

func getUsersSignIn(r render.Render) {
	r.HTML(http.StatusOK, "users/sign-in", nil)
}

func getUsersUnlock(r render.Render) {
	r.HTML(http.StatusOK, "users/unlock", nil)
}
