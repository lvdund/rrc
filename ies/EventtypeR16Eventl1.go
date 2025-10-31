package ies

// EventType-r16-eventL1 ::= SEQUENCE
type EventtypeR16Eventl1 struct {
	L1Threshold   Meastriggerquantity
	Hysteresis    Hysteresis
	Timetotrigger Timetotrigger
}
