package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/scys12/simple-api-go/auth"
	"github.com/scys12/simple-api-go/controllers"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	register("GET", "/books", controllers.AllBooks, auth.TokenMiddleware)
	register("GET", "/books/{id}", controllers.FindBook, auth.TokenMiddleware)
	register("POST", "/books", controllers.CreateBook, auth.TokenMiddleware)
	register("PUT", "/books", controllers.UpdateBook, auth.TokenMiddleware)
	register("DELETE", "/books/{id}", controllers.DeleteBook, auth.TokenMiddleware)

	register("POST", "/user/register", controllers.Register, nil)
	register("POST", "/user/login", controllers.Login, nil)
}

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		r := router.Methods(route.Method).
			Path(route.Pattern)
		if route.Middleware != nil {
			r.Handler(route.Middleware(route.Handler))
		} else {
			r.Handler(route.Handler)
		}
	}
	return router
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
