package main

import (
    "fmt"
    "net"
    "net/http"
    "time"

    "SimpleMVC/app/controller"
    "SimpleMVC/app/service"

    "github.com/gorilla/mux"
    "github.com/thoas/go-funk"
)

var routing *service.Routing

func init() {
    service.InitContainer()
    service.Container.GetLogger().App.Println("Initializing server")
    routing = service.Container.GetRouting()
}

func main() {
    routing.AddController(&controller.IndexController{})
    http.Handle("/", routing.RouteHandler())

    address := fmt.Sprintf("%v:%v", service.Container.GetConfig().Get("server.host"), service.Container.GetConfig().Get("server.port"))

    server := &http.Server{
        Addr:         address,
        ReadTimeout:  service.Container.GetConfig().GetDuration("server.timeout.read") * time.Second,
        WriteTimeout: service.Container.GetConfig().GetDuration("server.timeout.write") * time.Second,
        IdleTimeout:  service.Container.GetConfig().GetDuration("server.timeout.idle") * time.Second,
        ErrorLog:     service.Container.GetLogger().App,
    }
    service.Container.GetLogger().App.Println(fmt.Sprintf("Server started at http://%s", address))

    listener, err := net.Listen("tcp", address)
    if err != nil {
        service.Container.GetLogger().App.Fatal(err)
    }

    go func() {
        err := server.Serve(listener)
        if err != nil {
            service.Container.GetLogger().App.Fatal(err)
        }
    }()
    defer func() {_ = server.Close()}()

    var command string
    service.Container.GetLogger().App.Println(`Enter "exit", "quit" or "q" for closing application.`)

    for funk.Contains([]string{"exit", "quit", "q"}, command) == false {
        _, _ = fmt.Scanln(&command)
        switch command {
            case "exit", "quit", "q":
                service.Container.GetLogger().App.Println("Bye Bye...")
            case "routing":
                service.Container.GetLogger().App.Println("Project routes:")
                _ = routing.RouteHandler().Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
                    path, err := route.GetPathTemplate()
                    if err != nil {
                        return err
                    }
                    fmt.Println(fmt.Sprintf("Name: %s, URI_TEMPLATE: %s", route.GetName(), path))

                    return nil
                })
        default:
            service.Container.GetLogger().App.Printf(`Command "%s" is unknouwn.`, command)
        }
    }
}
