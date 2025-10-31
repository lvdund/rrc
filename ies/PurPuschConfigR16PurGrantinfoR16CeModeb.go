package ies

import "rrc/utils"

// PUR-PUSCH-Config-r16-pur-GrantInfo-r16-ce-ModeB ::= SEQUENCE
type PurPuschConfigR16PurGrantinfoR16CeModeb struct {
	SubprbAllocationR16  utils.BOOLEAN
	NumrusR16            utils.BOOLEAN
	PrbAllocationinfoR16 utils.BITSTRING `lb:8,ub:8`
	McsR16               utils.BITSTRING `lb:4,ub:4`
	NumrepetitionsR16    utils.BITSTRING `lb:3,ub:3`
}
