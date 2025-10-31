package ies

import "rrc/utils"

// LAA-Parameters-r13-secondSlotStartingPosition-r13 ::= ENUMERATED
type LaaParametersR13SecondslotstartingpositionR13 struct {
	Value utils.ENUMERATED
}

const (
	LaaParametersR13SecondslotstartingpositionR13EnumeratedNothing = iota
	LaaParametersR13SecondslotstartingpositionR13EnumeratedSupported
)
