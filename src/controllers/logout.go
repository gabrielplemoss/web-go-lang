package controllers

import (
	"agenda-online/src/cookies"
	"net/http"
)

func FazerLogout(w http.ResponseWriter, r *http.Request) {
	cookies.Deletar(w)
	http.Redirect(w, r, "/", 302)
}
