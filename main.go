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
    service.InitLogger()
    service.InitConfig()
    service.Logger.App.Println("Initializing server")
    routing = service.GetRouting()
}

func main() {
    routing.AddController(&controller.IndexController{})
    http.Handle("/", routing.RouteHandler())

    address := fmt.Sprintf("%v:%v", service.Configuration.Get("server.host"), service.Configuration.Get("server.port"))

    server := &http.Server{
        Addr:         address,
        ReadTimeout:  service.Configuration.GetDuration("server.timeout.read") * time.Second,
        WriteTimeout: service.Configuration.GetDuration("server.timeout.write") * time.Second,
        IdleTimeout:  service.Configuration.GetDuration("server.timeout.idle") * time.Second,
        ErrorLog:     service.Logger.App,
    }
    service.Logger.App.Println(fmt.Sprintf("Server started at http://%s", address))

    listener, err := net.Listen("tcp", address)
    if err != nil {
        service.Logger.App.Fatal(err)
    }

    go func() {
        err := server.Serve(listener)
        if err != nil {
            service.Logger.App.Fatal(err)
        }
    }()
    defer func() {_ = server.Close()}()

    var command string
    service.Logger.App.Println(`Type "exit" for closing application`)

    for funk.Contains([]string{"exit", "quit", "q"}, command) == false {
        _, _ = fmt.Scanln(&command)
        switch command {
            case "exit", "quit", "q":
                service.Logger.App.Println("Bye Bye...")
            case "routing":
                service.Logger.App.Println("Project routes:")
                _ = routing.RouteHandler().Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
                    path, err := route.GetPathTemplate()
                    if err != nil {
                        return err
                    }
                    fmt.Println(fmt.Sprintf("Name: %s, URI_TEMPLATE: %s", route.GetName(), path))

                    return nil
                })
        default:
            service.Logger.App.Printf(`Command "%s" is unknouwn.`, command)
        }
    }
}
