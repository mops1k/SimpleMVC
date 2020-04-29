package service

import (
    "bytes"
    "fmt"
    "reflect"
    "strings"
    "text/template"

    "github.com/arthurkushman/pgo"
)

type Controller interface {
    Action(c *Context) string
    ConfigName() string
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
        Container.GetLogger().App.Panic(err)
    }

    writer := &bytes.Buffer{}

    err = goTemplate.Execute(writer, vars)
    if err != nil {
        Container.GetLogger().App.Panic(err)
    }

    return string(writer.Bytes())
}

func (bc *BaseController) RenderString(value interface{}) string {
    return fmt.Sprintf("%v", value)
}

func (bc *BaseController) ConfigName() string {
    return strings.Replace(reflect.TypeOf(bc).String(), "*", "", -1)
}

func (cc *ControllerCollection) Add(c Controller) *ControllerCollection {
    if pgo.InArray(c, cc.controllers) == false {
        cc.controllers = append(cc.controllers, c)
    }

    return cc
}

func (cc *ControllerCollection) GetAll() []Controller {
    return cc.controllers
}
