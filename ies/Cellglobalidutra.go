package ies

import "rrc/utils"

// CellGlobalIdUTRA ::= SEQUENCE
type Cellglobalidutra struct {
	PlmnIdentity PlmnIdentity
	Cellidentity utils.BITSTRING
}
