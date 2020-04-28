package http

import (
    "net/http"

    "github.com/gookit/event"
)

type OnRequestEvent struct {
    event.BasicEvent
    request *http.Request
}

const OnRequest = "http.on_request"

func NewOnRequestEvent(r *http.Request) *OnRequestEvent {
    return &OnRequestEvent{request: r}
}

func (e *OnRequestEvent) Request() *http.Request {
    return e.request
}

func (e *OnRequestEvent) Name() string {
    return OnRequest
}

