package middlewares

import (
	"agenda-online/src/autenticacao"
	"agenda-online/src/cookies"
	"agenda-online/src/respostas"
	"net/http"
)

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		valores, err := cookies.LerCookie(r)
		r.Header.Add("Authorization", "Bearer "+valores["token"])
		if err != nil {
			http.Redirect(w, r, "/", 302)
			return
		}

		if erro := autenticacao.ValidarToken(r); erro != nil {
			respostas.Erro(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
