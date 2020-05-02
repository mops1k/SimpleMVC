package service

import (
    "sync"

    "github.com/CloudyKit/jet"
)

type container struct {
    database *Database
    config *Config
    logger *Log
    routing *Routing
    controllerCollection *ControllerCollection
    template *template
    lock *sync.Mutex
}

var Container *container

func InitContainer() {
    if Container == nil {
        Lock.Lock()
        defer Lock.Unlock()
        Container = &container{lock: Lock}
    }
}

func (c *container) GetDatabase() *Database {
    if c.database == nil {
        c.lock.Lock()
        defer c.lock.Unlock()
        c.database = &Database{}
        c.database.SetDialect(c.GetConfig().GetString("database.type"))
        c.database.SetUrl(c.GetConfig().GetString("database.url"))
        c.database.SetLogger(c.GetLogger().Database)
    }

    return c.database
}

func (c *container) GetConfig() *Config {
    if c.config == nil {
        c.lock.Lock()
        defer c.lock.Unlock()
        c.config = initConfig()
    }

    return c.config
}

func (c *container) GetLogger() *Log {
    if c.logger == nil {
        c.lock.Lock()
        defer c.lock.Unlock()
        c.logger = initLogger()
    }

    return c.logger
}

func (c *container) GetRouting() *Routing {
    if c.routing == nil {
        c.lock.Lock()
        defer c.lock.Unlock()
        c.routing = initRouter()
    }

    return c.routing
}

func (c *container) GetControllerCollection() *ControllerCollection {
    if c.controllerCollection == nil {
        c.lock.Lock()
        defer c.lock.Unlock()
        c.controllerCollection = &ControllerCollection{}
    }

    return c.controllerCollection
}

func (c *container) GetTemplate() *template {
    if c.template == nil {
        c.lock.Lock()
        defer c.lock.Unlock()
        c.template = &template{view: jet.NewHTMLSet("./templates/")}
    }

    return c.template
}
