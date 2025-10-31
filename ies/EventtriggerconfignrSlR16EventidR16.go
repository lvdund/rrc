package ies

// EventTriggerConfigNR-SL-r16-eventId-r16 ::= CHOICE
// Extensible
const (
	EventtriggerconfignrSlR16EventidR16ChoiceNothing = iota
	EventtriggerconfignrSlR16EventidR16ChoiceEventc1
	EventtriggerconfignrSlR16EventidR16ChoiceEventc2R16
)

type EventtriggerconfignrSlR16EventidR16 struct {
	Choice     uint64
	Eventc1    *EventtriggerconfignrSlR16EventidR16Eventc1
	Eventc2R16 *EventtriggerconfignrSlR16EventidR16Eventc2R16
}
