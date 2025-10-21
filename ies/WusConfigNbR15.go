package ies

import "rrc/utils"

// WUS-Config-NB-r15 ::= SEQUENCE
// Extensible
type WusConfigNbR15 struct {
	MaxdurationfactorR15   WusMaxdurationfactorNbR15
	NumposR15              utils.ENUMERATED
	NumdrxCyclesrelaxedR15 utils.ENUMERATED
	TimeoffsetdrxR15       utils.ENUMERATED
	TimeoffsetEdrxShortR15 utils.ENUMERATED
	TimeoffsetEdrxLongR15  *utils.ENUMERATED
}
