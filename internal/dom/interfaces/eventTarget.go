package main

type EventTarget interface {
	addEventListener(type_ string, callback *EventListener, options *AddEventListenerOptions)
	removeEventListener(type_ string, callback *EventListener, options *EventListenerOptions)
	dispatchEvent(event *Event) bool
}
