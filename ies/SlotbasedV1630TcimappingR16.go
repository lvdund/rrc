package ies

import "rrc/utils"

// SlotBased-v1630-tciMapping-r16 ::= ENUMERATED
type SlotbasedV1630TcimappingR16 struct {
	Value utils.ENUMERATED
}

const (
	SlotbasedV1630TcimappingR16EnumeratedNothing = iota
	SlotbasedV1630TcimappingR16EnumeratedCyclicmapping
	SlotbasedV1630TcimappingR16EnumeratedSequentialmapping
)
