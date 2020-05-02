package main

import (
    "fmt"
    "net"
    "net/http"
    "os"
    "time"

    "github.com/c-bata/go-prompt"

    "SimpleMVC/app/command"
    "SimpleMVC/app/service"
    "SimpleMVC/config"
)

var routing *service.Routing
var cc = &service.CommandCollection{}

func init() {
    service.InitContainer()
    _ = service.Container.GetLogger().App.Info("Initializing application")
    routing = service.Container.GetRouting()
}

func main() {
    if service.Container.GetConfig().GetBool("database.enabled") {
        db := service.Container.GetDatabase()
        db.Connect()
        defer db.Close()
    }

    config.InitControllers()
    routing.HandleControllers()
    http.Handle("/", routing.RouteHandler())

    address := fmt.Sprintf("%v:%v", service.Container.GetConfig().Get("server.host"), service.Container.GetConfig().Get("server.port"))

    server := &http.Server{
        Addr:         address,
        ReadTimeout:  service.Container.GetConfig().GetDuration("server.timeout.read") * time.Second,
        WriteTimeout: service.Container.GetConfig().GetDuration("server.timeout.write") * time.Second,
        IdleTimeout:  service.Container.GetConfig().GetDuration("server.timeout.idle") * time.Second,
        ErrorLog:     service.Container.GetLogger().Http,
    }
    _ = service.Container.GetLogger().App.Info(fmt.Sprintf("Server started at http://%s", address))

    listener, err := net.Listen("tcp", address)
    if err != nil {
        _ = service.Container.GetLogger().App.Critical(err.Error())
        os.Exit(1)
    }

    go func() {
        err := server.Serve(listener)
        if err != nil {
            _ = service.Container.GetLogger().App.Critical(err.Error())
            os.Exit(1)
        }
    }()
    defer func() { _ = server.Close() }()

    _ = service.Container.GetLogger().App.Info(`Enter "exit" for closing application.`)

    cc.Add(&command.ExitCommand{})
    p := prompt.New(executor, completer)
    p.Run()
}

func executor(t string) {
    c := cc.Get(t)
    c.Action()
}

func completer(t prompt.Document) []prompt.Suggest {
    var completer []prompt.Suggest
    for _, c := range cc.GetAll() {
        completer = append(completer, prompt.Suggest{Text: c.Name(), Description: c.Description()})
    }

    return completer
}
