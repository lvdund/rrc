package ies

import "rrc/utils"

// MeasParameters-v1310-whiteCellList-r13 ::= ENUMERATED
type MeasparametersV1310WhitecelllistR13 struct {
	Value utils.ENUMERATED
}

const (
	MeasparametersV1310WhitecelllistR13EnumeratedNothing = iota
	MeasparametersV1310WhitecelllistR13EnumeratedSupported
)
