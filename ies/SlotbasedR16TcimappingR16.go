package ies

import "rrc/utils"

// SlotBased-r16-tciMapping-r16 ::= ENUMERATED
type SlotbasedR16TcimappingR16 struct {
	Value utils.ENUMERATED
}

const (
	SlotbasedR16TcimappingR16EnumeratedNothing = iota
	SlotbasedR16TcimappingR16EnumeratedCyclicmapping
	SlotbasedR16TcimappingR16EnumeratedSequentialmapping
)
