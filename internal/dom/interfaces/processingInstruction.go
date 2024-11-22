package dom

// ProcessingInstruction Implementation
type ProcessingInstruction struct {
	CharacterData
	target string
}

type ProcessingInstructionInterface interface {
	CharacterDataInterface
}
