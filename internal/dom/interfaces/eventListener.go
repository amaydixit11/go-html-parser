package main

type EventListener interface {
	handleEvent(event *Event)
}
type EventListenerOptions struct {
	capture bool
}

type AddEventListenerOptions struct {
	EventListenerOptions
	passive bool
	once    bool
	signal  *AbortSignal
}
