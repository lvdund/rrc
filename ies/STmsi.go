package ies

import "rrc/utils"

// S-TMSI ::= SEQUENCE
type STmsi struct {
	Mmec  Mmec
	MTmsi utils.BITSTRING `lb:32,ub:32`
}
