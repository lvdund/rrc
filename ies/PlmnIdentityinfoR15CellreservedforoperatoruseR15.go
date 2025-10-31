package ies

import "rrc/utils"

// PLMN-IdentityInfo-r15-cellReservedForOperatorUse-r15 ::= ENUMERATED
type PlmnIdentityinfoR15CellreservedforoperatoruseR15 struct {
	Value utils.ENUMERATED
}

const (
	PlmnIdentityinfoR15CellreservedforoperatoruseR15EnumeratedNothing = iota
	PlmnIdentityinfoR15CellreservedforoperatoruseR15EnumeratedReserved
	PlmnIdentityinfoR15CellreservedforoperatoruseR15EnumeratedNotreserved
)
