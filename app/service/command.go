package service

type Command interface {
    Name() string
    Description() string
    Action()
}

type CommandCollection struct {
    collection map[string]Command
}

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
