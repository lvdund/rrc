package ies

import "rrc/utils"

// PLMN-IdentityInfo-r15-cellReservedForOperatorUse-CRS-r15 ::= ENUMERATED
type PlmnIdentityinfoR15CellreservedforoperatoruseCrsR15 struct {
	Value utils.ENUMERATED
}

const (
	PlmnIdentityinfoR15CellreservedforoperatoruseCrsR15EnumeratedNothing = iota
	PlmnIdentityinfoR15CellreservedforoperatoruseCrsR15EnumeratedReserved
	PlmnIdentityinfoR15CellreservedforoperatoruseCrsR15EnumeratedNotreserved
)
