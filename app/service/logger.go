package service

import (
    "log"
    "os"
)

type Log struct {
    App *log.Logger
}

func initLogger() *Log {
    return &Log{
            App: log.New(os.Stdout, "[app] ", log.LstdFlags),
        }
}
