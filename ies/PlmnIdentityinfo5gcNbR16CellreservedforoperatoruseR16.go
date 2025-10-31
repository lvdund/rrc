package ies

import "rrc/utils"

// PLMN-IdentityInfo-5GC-NB-r16-cellReservedForOperatorUse-r16 ::= ENUMERATED
type PlmnIdentityinfo5gcNbR16CellreservedforoperatoruseR16 struct {
	Value utils.ENUMERATED
}

const (
	PlmnIdentityinfo5gcNbR16CellreservedforoperatoruseR16EnumeratedNothing = iota
	PlmnIdentityinfo5gcNbR16CellreservedforoperatoruseR16EnumeratedReserved
	PlmnIdentityinfo5gcNbR16CellreservedforoperatoruseR16EnumeratedNotreserved
)
