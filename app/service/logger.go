package service

import (
    "log"
    "os"
)

type Log struct {
    App *log.Logger
}

var Logger *Log

func InitLogger() *Log {
    Lock.Lock()
    defer Lock.Unlock()

    if Logger == nil {
        Logger = &Log{
            App: log.New(os.Stdout, "[app] ", log.LstdFlags),
        }
    }

    return Logger
}
