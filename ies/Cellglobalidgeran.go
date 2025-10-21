package ies

import "rrc/utils"

// CellGlobalIdGERAN ::= SEQUENCE
type Cellglobalidgeran struct {
	PlmnIdentity     PlmnIdentity
	Locationareacode utils.BITSTRING
	Cellidentity     utils.BITSTRING
}
