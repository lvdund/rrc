package ies

import "rrc/utils"

// PLMN-IdentityInfo-NB-r13-cellReservedForOperatorUse-r13 ::= ENUMERATED
type PlmnIdentityinfoNbR13CellreservedforoperatoruseR13 struct {
	Value utils.ENUMERATED
}

const (
	PlmnIdentityinfoNbR13CellreservedforoperatoruseR13EnumeratedNothing = iota
	PlmnIdentityinfoNbR13CellreservedforoperatoruseR13EnumeratedReserved
	PlmnIdentityinfoNbR13CellreservedforoperatoruseR13EnumeratedNotreserved
)
