package service

import (
    "fmt"
    "net/http"
    "time"

    "github.com/gookit/event"
    "github.com/gorilla/mux"

    httpEvent "SimpleMVC/app/service/event/http"
)

type Routing struct {
    router *mux.Router
}

func initRouter() *Routing {
    router := &Routing{router: mux.NewRouter()}
    staticDirName := "static"
    router.router.PathPrefix("/" + staticDirName).
        Handler(http.StripPrefix(
            "/" + staticDirName + "/",
            http.FileServer(http.Dir("/" + staticDirName + "/")))).Name(staticDirName)

    return router
}

func (r *Routing) addController(c Controller, methods ...string) {
    methods = r.setDefaultMethods(methods)
    pathName, path := c.Name(), Container.GetConfig().GetString(c.Name()+".path")
    r.router.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
        onRequestEvent := httpEvent.NewOnRequestEvent(request)
        event.AddEvent(onRequestEvent)

        _ = event.FireEvent(onRequestEvent)
        start := time.Now()
        var context = &Context{response: writer, request: request, statusCode: http.StatusOK}
        content := c.Action(context)
        onResponseEvent := httpEvent.NewOnResponseEvent(request.Response, content)
        event.AddEvent(onResponseEvent)

        if context.headers != nil {
            for name, value := range context.headers {
                writer.Header().Add(name, value)
            }
        }
        writer.WriteHeader(context.statusCode)

        _ = event.FireEvent(onResponseEvent)
        _, err := fmt.Fprint(writer, content)
        if err != nil {
            Container.GetLogger().App.Fatal(err)
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

    Container.GetLogger().App.Println(
        req.RequestURI,
        req.Method,
        requesterIP,
        time.Since(start),
        req.UserAgent(),
    )
}

func (r *Routing) HandleControllers() {
    for _, controller := range Container.GetControllerCollection().GetAll() {
        r.addController(controller, Container.GetConfig().GetStringSlice(controller.Name()+".methods")...)
    }
}
