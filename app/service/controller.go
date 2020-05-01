package service

import (
    "bytes"
    "fmt"
    "os"
    "text/template"

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

func (bc *BaseController) Render(vars map[interface{}]interface{}, filenames ...string) string {
    var paths []string
    for _, filename := range filenames {
        paths = append(paths, "templates/"+filename)
    }
    goTemplate, err := template.ParseFiles(paths...)
    if err != nil {
        Container.GetLogger().App.Critical(err.Error())
        os.Exit(2)
    }

    writer := &bytes.Buffer{}

    err = goTemplate.Execute(writer, vars)
    if err != nil {
        Container.GetLogger().App.Critical(err.Error())
        os.Exit(2)
    }

    return writer.String()
}

func (bc *BaseController) RenderString(value interface{}) string {
    return fmt.Sprintf("%v", value)
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
