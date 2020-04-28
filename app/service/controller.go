package service

import (
    "bytes"
    "fmt"
    "text/template"
)

type Controller interface {
    Action(c *Context) string
    Name() (string, string)
}

type BaseController struct {
    Controller
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
