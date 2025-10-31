package ies

import "rrc/utils"

// PUCCH-Grp-CarrierTypes-r16-fr2-r16 ::= ENUMERATED
type PucchGrpCarriertypesR16Fr2R16 struct {
	Value utils.ENUMERATED
}

const (
	PucchGrpCarriertypesR16Fr2R16EnumeratedNothing = iota
	PucchGrpCarriertypesR16Fr2R16EnumeratedSupported
)
