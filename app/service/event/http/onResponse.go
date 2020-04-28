package http

import (
    "net/http"

    "github.com/gookit/event"
)

type OnResponseEvent struct {
    event.BasicEvent
    response *http.Response
    content string
}

const OnResponse = "http.on_response"

func NewOnResponseEvent(response *http.Response, content string) *OnResponseEvent {
    return &OnResponseEvent{response: response, content: content}
}

func (e *OnResponseEvent) Response() *http.Response {
    return e.response
}

func (e *OnResponseEvent) Name() string {
    return OnResponse
}
