package main

import (
    "fmt"
    "net"
    "net/http"
    "time"

    "SimpleMVC/app/controller"
    "SimpleMVC/app/service"

    "github.com/arthurkushman/pgo"
    "github.com/gorilla/mux"
)

var routing *service.Routing

func init() {
    service.InitContainer()
    service.Container.GetLogger().App.Println("Initializing application")
    routing = service.Container.GetRouting()
}

func main() {
    if service.Container.GetConfig().GetBool("database.enabled") {
        db := service.Container.GetDatabase()
        db.Connect()
        defer db.Close()
    }

    service.Container.GetControllerCollection().
        Add(&controller.IndexController{})

    routing.HandleControllers()

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

    for pgo.InArray(command, []string{"exit", "quit", "q"}) == false {
        _, _ = fmt.Scanln(&command)
        switch command {
            case "exit", "quit", "q":
                service.Container.GetLogger().App.Println("Bye Bye...")
            case "routing":
                service.Container.GetLogger().App.Println("Project routes:")
                _ = routing.RouteHandler().Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
                    path, err := route.GetPathTemplate()
                    methods, _ := route.GetMethods()
                    if err != nil {
                        return err
                    }
                    fmt.Println(fmt.Sprintf("Name: %s, URI_TEMPLATE: %s METHODS: %s", route.GetName(), path, methods))

                    return nil
                })
        default:
            service.Container.GetLogger().App.Printf(`Command "%s" is unknouwn.`, command)
        }
    }
}
