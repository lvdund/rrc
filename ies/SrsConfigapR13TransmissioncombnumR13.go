package ies

import "rrc/utils"

// SRS-ConfigAp-r13-transmissionCombNum-r13 ::= ENUMERATED
type SrsConfigapR13TransmissioncombnumR13 struct {
	Value utils.ENUMERATED
}

const (
	SrsConfigapR13TransmissioncombnumR13EnumeratedNothing = iota
	SrsConfigapR13TransmissioncombnumR13EnumeratedN2
	SrsConfigapR13TransmissioncombnumR13EnumeratedN4
)
