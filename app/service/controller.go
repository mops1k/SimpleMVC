package service

import (
    "github.com/arthurkushman/pgo"
)

type Controller interface {
    Action(c *Context) string
    Name() string
}

type BaseController struct {
    Controller
}

type ControllerCollection struct {
    controllers []Controller
}

func (bc *BaseController) Render(filename string, vars map[string]interface{}) string {
    return Container.GetTemplate().RenderTemplate(filename, vars)
}

func (bc *BaseController) RenderString(content string, vars map[string]interface{}) string {
    return Container.GetTemplate().RenderString(content, vars)
}

func (cc *ControllerCollection) Add(c Controller) *ControllerCollection {
    if !pgo.InArray(c, cc.controllers) {
        cc.controllers = append(cc.controllers, c)
    }

    return cc
}

func (cc *ControllerCollection) GetAll() []Controller {
    return cc.controllers
}
