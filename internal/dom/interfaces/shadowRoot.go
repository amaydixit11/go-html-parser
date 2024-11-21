package dom

type ShadowRootMode string

const (
	open   ShadowRootMode = "open"
	closed ShadowRootMode = "closed"
)

type SlotAssignmentMode string

const (
	manual SlotAssignmentMode = "manual"
	named  SlotAssignmentMode = "named"
)

type ShadowRoot struct {
	Mode           ShadowRootMode
	DelegatesFocus bool
	SlotAssignment SlotAssignmentMode
	Clonable       bool
	Serializable   bool
	Host           *Element
	OnSlotChange   *EventHandler
}

func ShadowRootInit(mode ShadowRootMode) ShadowRoot {
	return ShadowRoot{
		Mode:           mode,
		DelegatesFocus: false,
		SlotAssignment: "named",
		Clonable:       false,
		Serializable:   false,
	}
}
