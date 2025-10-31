package ies

import "rrc/utils"

// SIB1-v1700-IEs-cellBarredNTN-r17 ::= ENUMERATED
type Sib1V1700IesCellbarredntnR17 struct {
	Value utils.ENUMERATED
}

const (
	Sib1V1700IesCellbarredntnR17EnumeratedNothing = iota
	Sib1V1700IesCellbarredntnR17EnumeratedBarred
	Sib1V1700IesCellbarredntnR17EnumeratedNotbarred
)
