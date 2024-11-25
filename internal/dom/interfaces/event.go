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

func (e *Event) GetType_() string {
	return e.type_
}

func (e *Event) GetTarget() *EventTarget {
	return e.target
}

func (e *Event) GetSrcElement() *EventTarget {
	return e.srcElement
}

func (e *Event) GetCurrentTarget() *EventTarget {
	return e.currentTarget
}

func (e *Event) GetEventPhase() EventPhaseType {
	return e.eventPhase
}

func (e *Event) GetBubbles() bool {
	return e.bubbles
}

func (e *Event) GetCancelable() bool {
	return e.cancelable
}

func (e *Event) GetReturnValue() bool {
	return e.returnValue
}

func (e *Event) GetDefaultPrevented() bool {
	return e.defaultPrevented
}

func (e *Event) GetComposed() bool {
	return e.composed
}

func (e *Event) IsTrusted() bool {
	return e.isTrusted
}

func (e *Event) GetTimeStamp() DOMHighResTimeStamp {
	return e.timeStamp
}

func (e *Event) SetReturnValue(returnValue bool) bool {
	e.returnValue = returnValue
	return e.returnValue
}

func (e *Event) PreventDefault() {
	if e.cancelable {
		e.defaultPrevented = true
	}
}

func (e *Event) InitEvent(type_ string, bubbles bool, cancelable bool) {
	e.type_ = type_
	e.bubbles = bubbles
	e.cancelable = cancelable
	e.composed = false
	e.timeStamp = DOMHighResTimeStamp(0) // Can set actual timestamp if needed
	e.isTrusted = false                  // Event created via dispatchEvent, not a user agent event
}

func (e *Event) GetComposedPath() []*EventTarget {
	var composedPath []*EventTarget
	path := e.ComposedPath // You can expand this path according to the specification

	if len(path) == 0 {
		return composedPath
	}

	currentTarget := e.GetCurrentTarget()
	composedPath = append(composedPath, currentTarget)

	// Additional logic for handling closed-tree shadow roots, etc. (omitted for brevity)
	// This would loop through the path and adjust based on shadow DOM tree status.

	return composedPath
}
