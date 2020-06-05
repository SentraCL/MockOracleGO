package mocksentra

/*
 * BFF Mi Pago
 *
 * BFF de Mi Pago
 *
 * API version: 1.0-beta
 * Contact: evillagran@bancochile.cl
  * SEnTRA 2020
*/

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	//OK
	Route{
		"ActualizaPerfilUsingPOST",
		strings.ToUpper("Post"),
		"/v1/clientes/persona/perfil/",
		ActualizaPerfilUsingPOST,
	},
	//OK
	Route{
		"ActualizarAvatarUsingPOST",
		strings.ToUpper("Post"),
		"/v1/clientes/persona/perfil/avatar",
		ActualizarAvatarUsingPOST,
	},
	//OK
	Route{
		"ObtenerAvatarUsingGET",
		strings.ToUpper("Get"),
		"/v1/clientes/persona/perfil/avatar/{rut_cifrado}",
		ObtenerAvatarUsingGET,
	},
	//OK
	Route{
		"ObtenerPerfilUsingGET",
		strings.ToUpper("Get"),
		"/v1/clientes/persona/perfil/",
		ObtenerPerfilUsingGET,
	},
	//OK productos
	Route{
		"GetProductosClienteUsingGET",
		strings.ToUpper("Get"),
		"/v1/productos/{codigos}",
		GetProductosClienteUsingGET,
	},
	// ???
	Route{
		"DesuscribirUsingDELETE",
		strings.ToUpper("Delete"),
		"/public/v1/utilidades/notificaciones/push/{token_push}",
		DesuscribirUsingDELETE,
	},
	//???
	Route{
		"SuscribirUsingPOST",
		strings.ToUpper("Post"),
		"/v1/utilidades/notificaciones/push/",
		SuscribirUsingPOST,
	},
}
