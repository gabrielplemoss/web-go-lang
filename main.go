package main

import (
	"agenda-online/src/config"
	"agenda-online/src/cookies"
	"agenda-online/src/routes"
	"agenda-online/src/utils"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.Carregar()
	cookies.Configurar()
	utils.CarregarTemplates()
	rotas := routes.Configurar(mux.NewRouter())

	fileServer := http.FileServer(http.Dir("./assets/"))
	rotas.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	fmt.Println("rodadno 5000")
	http.ListenAndServe(fmt.Sprintf(":%s", config.Porta), rotas)

}
