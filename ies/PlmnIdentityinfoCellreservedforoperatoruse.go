package ies

import "rrc/utils"

// PLMN-IdentityInfo-cellReservedForOperatorUse ::= ENUMERATED
type PlmnIdentityinfoCellreservedforoperatoruse struct {
	Value utils.ENUMERATED
}

const (
	PlmnIdentityinfoCellreservedforoperatoruseEnumeratedNothing = iota
	PlmnIdentityinfoCellreservedforoperatoruseEnumeratedReserved
	PlmnIdentityinfoCellreservedforoperatoruseEnumeratedNotreserved
)
