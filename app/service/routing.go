package service

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sync"
	"time"
)

type Routing struct {
	router *mux.Router
}

var Router *Routing
var once sync.Once

func GetRouting() *Routing {
	if Router == nil {
		once.Do(func() {
			Router = &Routing{router: mux.NewRouter()}
			Router.router.PathPrefix("/static/").
				Handler(http.StripPrefix(
					"/static/",
					http.FileServer(http.Dir("/static/")))).Name("static")
		})
	}

	return Router
}

func (r *Routing) AddController(c Controller, methods ...string) {
	methods = r.setDefaultMethods(methods)
	pathName, path := c.Name()
	r.router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		start := time.Now()
		var context = &Context{response: writer, request: request, statusCode: http.StatusOK}
		content := c.Action(context)

		if context.headers != nil {
			for name, value := range context.headers {
				writer.Header().Add(name, value)
			}
		}
		writer.WriteHeader(context.statusCode)

		_, err := fmt.Fprint(writer, content)
		if err != nil {
			Logger.App.Fatal(err)
		}

		r.logRequest(start, request)
	}).Methods(methods...).Name(pathName)
}

func (r *Routing) RouteHandler() *mux.Router {
	return r.router
}

func (r *Routing) setDefaultMethods(methods []string) []string {
	if len(methods) == 0 {
		methods = append(methods, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete)
	}

	return methods
}

func (r *Routing) logRequest(start time.Time, req *http.Request) {
	requesterIP := req.RemoteAddr

	Logger.App.Println(
		req.RequestURI,
		req.Method,
		requesterIP,
		time.Since(start),
		req.UserAgent(),
	)
}
