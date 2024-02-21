package main

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type FollowRequestSent struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type EventsCounter interface {
	CountEvent() error
}

func NewFollowRequestSentHandler(counter EventsCounter) cqrs.EventHandler {
	return &FollowRequestSentHandler{
		counter: counter,
	}
}

type FollowRequestSentHandler struct {
	counter EventsCounter
}

func (h *FollowRequestSentHandler) HandlerName() string {
	return "FollowRequestSentHandler"
}

func (h *FollowRequestSentHandler) NewEvent() interface{} {
	return &FollowRequestSent{}
}

func (h *FollowRequestSentHandler) Handle(ctx context.Context, event any) error {
	//e := event.(*FollowRequestSent)

	return h.counter.CountEvent()
}
