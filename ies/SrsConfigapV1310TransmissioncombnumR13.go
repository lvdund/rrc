package ies

import "rrc/utils"

// SRS-ConfigAp-v1310-transmissionCombNum-r13 ::= ENUMERATED
type SrsConfigapV1310TransmissioncombnumR13 struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigapV1310TransmissioncombnumR13EnumeratedNothing = iota
	SrsConfigapV1310TransmissioncombnumR13EnumeratedN2
	SrsConfigapV1310TransmissioncombnumR13EnumeratedN4
)
