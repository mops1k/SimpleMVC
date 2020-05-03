package service

import "SimpleMVC/app/service/command"

type Command interface {
    Name() string
    Description() string
    Action(ctx command.Context)
}

type CommandCollection struct {
    collection map[string]Command
}

var Commands *CommandCollection

func (cc *CommandCollection) Add(c Command) *CommandCollection {
    if cc.collection == nil {
        cc.collection = make(map[string]Command)
    }

    cc.collection[c.Name()] = c

    return cc
}

func (cc *CommandCollection) Get(name string) Command {
    return cc.collection[name]
}

func (cc *CommandCollection) GetAll() map[string]Command {
    return cc.collection
}

func (cc *CommandCollection) Has(name string) bool {
    if cc.collection[name] == nil {
        return false
    }

    return true
}

func init() {
    Commands = &CommandCollection{}
}
