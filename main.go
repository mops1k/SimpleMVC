package main

import (
	"SimpleMVC/app/controller"
	"SimpleMVC/app/service"
	"fmt"
	"net/http"
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
		Handler:     routing.RouteHandler(),
		Addr:        address,
		IdleTimeout: 30,
	}
	service.Logger.App.Println(fmt.Sprintf("Server started at http://%s", address))
	service.Logger.App.Panic(server.ListenAndServe())
}
