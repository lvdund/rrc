package ies

import "rrc/utils"

// MeasParameters-v1310-extendedMaxObjectId-r13 ::= ENUMERATED
type MeasparametersV1310ExtendedmaxobjectidR13 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1310ExtendedmaxobjectidR13EnumeratedNothing = iota
	MeasparametersV1310ExtendedmaxobjectidR13EnumeratedSupported
)
