package config

import (
    "SimpleMVC/app/controller"
    "SimpleMVC/app/service"
)

var collection *service.ControllerCollection

func InitControllers() {
    collection = service.Container.GetControllerCollection()

    // Add your controllers here
    collection.Add(&controller.IndexController{})
}
