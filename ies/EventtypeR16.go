package ies

// EventType-r16 ::= CHOICE
// Extensible
const (
	EventtypeR16ChoiceNothing = iota
	EventtypeR16ChoiceOutofcoverage
	EventtypeR16ChoiceEventl1
)

type EventtypeR16 struct {
	Choice        uint64
	Outofcoverage *struct{}
	Eventl1       *EventtypeR16Eventl1
}
