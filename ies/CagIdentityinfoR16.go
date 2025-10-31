package ies

import "rrc/utils"

// CAG-IdentityInfo-r16 ::= SEQUENCE
type CagIdentityinfoR16 struct {
	CagIdentityR16               utils.BITSTRING `lb:32,ub:32`
	ManualcagselectionallowedR16 *utils.BOOLEAN
}
