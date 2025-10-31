package ies

import "rrc/utils"

// PLMN-IdentityInfo-v1530-cellReservedForOperatorUse-CRS-r15 ::= ENUMERATED
type PlmnIdentityinfoV1530CellreservedforoperatoruseCrsR15 struct {
	Value utils.ENUMERATED
}

const (
	PlmnIdentityinfoV1530CellreservedforoperatoruseCrsR15EnumeratedNothing = iota
	PlmnIdentityinfoV1530CellreservedforoperatoruseCrsR15EnumeratedReserved
	PlmnIdentityinfoV1530CellreservedforoperatoruseCrsR15EnumeratedNotreserved
)
