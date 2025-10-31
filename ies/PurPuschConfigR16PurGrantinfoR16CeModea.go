package ies

import "rrc/utils"

// PUR-PUSCH-Config-r16-pur-GrantInfo-r16-ce-ModeA ::= SEQUENCE
type PurPuschConfigR16PurGrantinfoR16CeModea struct {
	NumrusR16            utils.BITSTRING `lb:2,ub:2`
	PrbAllocationinfoR16 utils.BITSTRING `lb:10,ub:10`
	McsR16               utils.BITSTRING `lb:4,ub:4`
	NumrepetitionsR16    utils.BITSTRING `lb:3,ub:3`
}
