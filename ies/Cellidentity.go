package ies

import "rrc/utils"

// CellIdentity ::= BIT STRING (SIZE (28))
type Cellidentity struct {
	Value utils.BITSTRING `lb:28,ub:28`
}
