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
	if Logger == nil {
		Logger = &Log{
			App: log.New(os.Stdout, "[app] ", log.LstdFlags),
		}
	}

	return Logger
}
