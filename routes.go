package main

import "github.com/gorilla/mux"
import "net/http"

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)

	}

	return router
}

var routes = Routes{
	Route{
		"LegajoList",
		"GET",
		"/api/legajo/legajos",
		LegajoList,
	},
	Route{
		"LegajoShow",
		"GET",
		"/api/legajo/legajos/{id}",
		LegajoShow,
	},
	Route{
		"LegajoAdd",
		"POST",
		"/api/legajo/legajos",
		LegajoAdd,
	},
	Route{
		"LegajoUpdate",
		"PUT",
		"/api/legajo/legajos/{id}",
		LegajoUpdate,
	},
	Route{
		"LegajoRemove",
		"DELETE",
		"/api/legajo/legajos/{id}",
		LegajoRemove,
	},
	Route{
		"LegajosRemoveMasivo",
		"DELETE",
		"/api/legajo/legajos",
		LegajosRemoveMasivo,
	},
	Route{
		"Healthy",
		"GET",
		"/api/legajo/healthy",
		Healthy,
	},
}
