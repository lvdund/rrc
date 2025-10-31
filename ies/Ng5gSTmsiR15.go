package ies

import "rrc/utils"

// NG-5G-S-TMSI-r15 ::= BIT STRING (SIZE (48))
type Ng5gSTmsiR15 struct {
	Value utils.BITSTRING `lb:48,ub:48`
}
