package service

import (
    "bytes"
    "os"

    "github.com/CloudyKit/jet"
    "github.com/arthurkushman/pgo"
)

type template struct {
    view *jet.Set
}

func (t *template) AddFunc(key string, fn jet.Func) {
    t.view.AddGlobalFunc(key, fn)
}

func (t *template) AddGlobal(key string, value interface{}) {
    t.view.AddGlobal(key, value)
}

func (t *template) render(template *jet.Template, vars map[string]interface{}) string {
    varMap := make(jet.VarMap)
    for name, value := range vars {
        varMap.Set(name, value)
    }

    var w bytes.Buffer
    err := template.Execute(&w, varMap, nil)
    if err != nil {
        _ = Container.GetLogger().App.Critical(err.Error())
        os.Exit(1)
    }

    return w.String()
}

func (t *template) RenderTemplate(filename string, vars map[string]interface{}) string {
    template, err := t.view.GetTemplate(filename)
    if err != nil {
        _ = Container.GetLogger().App.Critical(err.Error())
        os.Exit(1)
    }

    return t.render(template, vars)
}

func (t *template) RenderString(content string, vars map[string]interface{}) string {
    template, err := t.view.Parse(pgo.Md5(content), content)
    if err != nil {
        _ = Container.GetLogger().App.Critical(err.Error())
        os.Exit(1)
    }

    return t.render(template, vars)
}
