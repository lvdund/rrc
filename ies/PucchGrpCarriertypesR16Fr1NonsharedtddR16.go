package ies

import "rrc/utils"

// PUCCH-Grp-CarrierTypes-r16-fr1-NonSharedTDD-r16 ::= ENUMERATED
type PucchGrpCarriertypesR16Fr1NonsharedtddR16 struct {
	Value utils.ENUMERATED
}

const (
	PucchGrpCarriertypesR16Fr1NonsharedtddR16EnumeratedNothing = iota
	PucchGrpCarriertypesR16Fr1NonsharedtddR16EnumeratedSupported
)
