package main

type AbortSignal struct {
	aborted bool
	reason  any
	onabort func(event *Event)
}

type AbortSignalInterface interface {
	EventTarget
	abort(reason any) *AbortSignal
	timeout(milliseconds uint64) *AbortSignal
	_any(signals []*AbortSignal) *AbortSignal
	throwIfAborted()
}
