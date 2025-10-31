package ies

import "rrc/utils"

// NG-5G-S-TMSI ::= BIT STRING (SIZE (48))
type Ng5gSTmsi struct {
	Value utils.BITSTRING `lb:48,ub:48`
}
