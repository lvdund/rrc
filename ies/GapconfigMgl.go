package ies

import "rrc/utils"

// GapConfig-mgl ::= ENUMERATED
type GapconfigMgl struct {
	Value utils.ENUMERATED
}

const (
	GapconfigMglEnumeratedNothing = iota
	GapconfigMglEnumeratedMs1dot5
	GapconfigMglEnumeratedMs3
	GapconfigMglEnumeratedMs3dot5
	GapconfigMglEnumeratedMs4
	GapconfigMglEnumeratedMs5dot5
	GapconfigMglEnumeratedMs6
)
