package ies

import "rrc/utils"

// SL-SourceIdentity-r17 ::= BIT STRING (SIZE (24))
type SlSourceidentityR17 struct {
	Value utils.BITSTRING `lb:24,ub:24`
}
