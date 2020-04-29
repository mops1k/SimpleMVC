package controller

import (
    "SimpleMVC/app/service"
)

type IndexController struct {
    *service.BaseController
}

func (ic *IndexController) Action(c *service.Context) string {
    return ic.RenderString("Welcome to SimpleMVC!")
}

func (ic *IndexController) Name() string {
    return "index"
}
