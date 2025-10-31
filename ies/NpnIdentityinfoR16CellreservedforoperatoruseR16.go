package ies

import "rrc/utils"

// NPN-IdentityInfo-r16-cellReservedForOperatorUse-r16 ::= ENUMERATED
type NpnIdentityinfoR16CellreservedforoperatoruseR16 struct {
	Value utils.ENUMERATED
}

const (
	NpnIdentityinfoR16CellreservedforoperatoruseR16EnumeratedNothing = iota
	NpnIdentityinfoR16CellreservedforoperatoruseR16EnumeratedReserved
	NpnIdentityinfoR16CellreservedforoperatoruseR16EnumeratedNotreserved
)
