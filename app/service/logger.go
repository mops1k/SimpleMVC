package service

import (
    "io"
    "log"
    "os"

    "github.com/lajosbencz/glo"
)

type Log struct {
    App  glo.Facility
    Http *log.Logger
    Database *log.Logger
}

func initLogger() *Log {
    l := &Log{}
    l.App = glo.NewStdFacility()
    l.Http = l.createFileLog("var/log/access.log", "[http] ")
    l.Database = l.createFileLog("var/log/database.log", "[database] ")
    return l
}

func (l *Log) createFileLog(filename string, prefix string) *log.Logger {
    file, err := os.OpenFile(filename, os.O_CREATE | os.O_APPEND, 0666)
    if err != nil {
        log.Panic(err)
    }

    writer := io.MultiWriter(file)

    return log.New(writer, prefix, log.LstdFlags)
}
