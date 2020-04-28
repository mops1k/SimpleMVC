package service

import (
	"log"
	"os"
	"sync"
)

type Log struct {
	App *log.Logger
}

var Logger *Log
var lock = &sync.Mutex{}

func InitLogger() *Log {
	lock.Lock()
	defer lock.Unlock()

	if Logger == nil {
		Logger = &Log{
			App: log.New(os.Stdout, "[app] ", log.LstdFlags),
		}
	}

	return Logger
}
