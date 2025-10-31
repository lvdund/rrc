package ies

import "rrc/utils"

// ReestabNCellInfo ::= SEQUENCE
type ReestabNCellInfo struct {
	CellIdentity  Cellidentity
	KeyGNodeBStar utils.BITSTRING `lb:0,ub:256`
	ShortMACI     ShortmacI
}
