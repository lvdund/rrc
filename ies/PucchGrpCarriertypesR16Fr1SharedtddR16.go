package ies

import "rrc/utils"

// PUCCH-Grp-CarrierTypes-r16-fr1-SharedTDD-r16 ::= ENUMERATED
type PucchGrpCarriertypesR16Fr1SharedtddR16 struct {
	Value utils.ENUMERATED
}

const (
	PucchGrpCarriertypesR16Fr1SharedtddR16EnumeratedNothing = iota
	PucchGrpCarriertypesR16Fr1SharedtddR16EnumeratedSupported
)
