package ies

import "rrc/utils"

// CellIdentity ::= BIT STRING (SIZE (36))
type Cellidentity struct {
	Value utils.BITSTRING `lb:36,ub:36`
}
