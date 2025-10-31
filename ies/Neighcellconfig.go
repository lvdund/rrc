package ies

import "rrc/utils"

// NeighCellConfig ::= BIT STRING (SIZE (2))
type Neighcellconfig struct {
	Value utils.BITSTRING `lb:2,ub:2`
}
