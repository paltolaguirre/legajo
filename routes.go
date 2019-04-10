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
		"/legajos",
		LegajoList,
	},
	Route{
		"LegajoShow",
		"GET",
		"/legajos/{id}",
		LegajoShow,
	},
	Route{
		"LegajoAdd",
		"POST",
		"/legajos",
		LegajoAdd,
	},
	Route{
		"LegajoUpdate",
		"PUT",
		"/legajos/{id}",
		LegajoUpdate,
	},
	Route{
		"LegajoPatch",
		"PATCH",
		"/legajos/{id}",
		LegajoPatch,
	},
	Route{
		"LegajoRemove",
		"DELETE",
		"/legajos/{id}",
		LegajoRemove,
	},
}
