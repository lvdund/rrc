package ies

import "rrc/utils"

// PUCCH-Grp-CarrierTypes-r16-fr1-NonSharedFDD-r16 ::= ENUMERATED
type PucchGrpCarriertypesR16Fr1NonsharedfddR16 struct {
	Value utils.ENUMERATED
}

const (
	PucchGrpCarriertypesR16Fr1NonsharedfddR16EnumeratedNothing = iota
	PucchGrpCarriertypesR16Fr1NonsharedfddR16EnumeratedSupported
)
