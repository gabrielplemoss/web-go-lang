package routes

import (
	"agenda-online/src/controllers"
	"agenda-online/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func Configurar(route *mux.Router) *mux.Router {
	route.HandleFunc("/", controllers.PaginaLogin).Methods(http.MethodGet)
	route.HandleFunc("/home", middlewares.Autenticar(controllers.Home)).Methods(http.MethodGet)
	route.HandleFunc("/login", controllers.LogarUsuario).Methods(http.MethodPost)
	route.HandleFunc("/cadastrar", controllers.CadastrarUsuario).Methods(http.MethodPost)

	route.HandleFunc("/contato", middlewares.Autenticar(controllers.CadastrarContato)).Methods(http.MethodPost)
	route.HandleFunc("/criar", middlewares.Autenticar(controllers.CadastrarContatoHtml)).Methods(http.MethodGet)
	route.HandleFunc("/contato/{contatoID}", middlewares.Autenticar(controllers.BuscarContato)).Methods(http.MethodGet)
	route.HandleFunc("/editar/contato/{contatoID}", middlewares.Autenticar(controllers.AtualizarContato)).Methods(http.MethodPost)
	route.HandleFunc("/deletar/contato/{contatoID}", middlewares.Autenticar(controllers.DeletarContato)).Methods(http.MethodGet)

	route.HandleFunc("/logout", middlewares.Autenticar(controllers.FazerLogout)).Methods(http.MethodGet)

	return route
}
