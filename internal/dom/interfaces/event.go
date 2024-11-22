package dom

type EventPhaseType uint8
type DOMHighResTimeStamp float64

const (
	NONE            EventPhaseType = 0
	CAPTURING_PHASE EventPhaseType = 1
	AT_TARGET       EventPhaseType = 2
	BUBBLING_PHASE  EventPhaseType = 3
)

type Event struct {
	type_            string
	target           *EventTarget
	srcElement       *EventTarget
	currentTarget    *EventTarget
	ComposedPath     []*EventTarget
	eventPhase       EventPhaseType
	bubbles          bool
	cancelable       bool
	returnValue      bool
	defaultPrevented bool
	composed         bool
	isTrusted        bool
	timeStamp        DOMHighResTimeStamp
}

type EventInterface interface {
	GetType_() string
	GetTarget() *EventTarget
	GetSrcElement() *EventTarget
	GetCurrentTarget() *EventTarget
	GetComposedPath() []*EventTarget
	GetEventPhase() EventPhaseType
	StopPropagation()
	StopImmediatePropagation()
	GetBubbles() bool
	GetCancelable() bool
	GetReturnValue() bool
	SetReturnValue(returnValue bool) bool
	PreventDefault()
	GetDefaultPrevented() bool
	GetComposed() bool
	IsTrusted() bool
	GetTimeStamp() DOMHighResTimeStamp
}

func EventInit(type_ string) Event {
	return Event{
		type_:      type_,
		bubbles:    false,
		cancelable: false,
		composed:   false,
	}
}

type EventTarget interface {
	addEventListener(type_ string, callback *EventListener, options *AddEventListenerOptions)
	removeEventListener(type_ string, callback *EventListener, options *EventListenerOptions)
	dispatchEvent(event *Event) bool
}

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

type AbortController struct {
	signal *AbortSignal
}

type AbortControllerInterface interface {
	Abort(reason any)
}

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
