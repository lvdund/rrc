package ies

import "rrc/utils"

// PhysCellIdGERAN ::= SEQUENCE
type Physcellidgeran struct {
	Networkcolourcode     utils.BITSTRING `lb:3,ub:3`
	Basestationcolourcode utils.BITSTRING `lb:3,ub:3`
}
