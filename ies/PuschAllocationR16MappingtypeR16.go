package ies

import "rrc/utils"

// PUSCH-Allocation-r16-mappingType-r16 ::= ENUMERATED
type PuschAllocationR16MappingtypeR16 struct {
	Value utils.ENUMERATED
}

const (
	PuschAllocationR16MappingtypeR16EnumeratedNothing = iota
	PuschAllocationR16MappingtypeR16EnumeratedTypea
	PuschAllocationR16MappingtypeR16EnumeratedTypeb
)
