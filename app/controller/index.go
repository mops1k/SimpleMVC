package controller

import (
    "SimpleMVC/app/service"
)

type IndexController struct {
    *service.BaseController
}

func (ic *IndexController) Action(c *service.Context) string {
    // return ic.RenderString(`Welcome to {{ Name }}!`, map[string]interface{}{"Name": "SimpleMVC"})
    return ic.Render("index.jet", map[string]interface{}{"Name": "SimpleMVC"})
}

func (ic *IndexController) Name() string {
    return "index"
}
