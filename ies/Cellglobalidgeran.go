package ies

import "rrc/utils"

// CellGlobalIdGERAN ::= SEQUENCE
type Cellglobalidgeran struct {
	PlmnIdentity     PlmnIdentity
	Locationareacode utils.BITSTRING `lb:16,ub:16`
	Cellidentity     utils.BITSTRING `lb:16,ub:16`
}
