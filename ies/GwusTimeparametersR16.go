package ies

import "rrc/utils"

// GWUS-TimeParameters-r16 ::= SEQUENCE
// Extensible
type GwusTimeparametersR16 struct {
	MaxdurationfactorR16   utils.ENUMERATED
	NumposR16              utils.ENUMERATED
	TimeoffsetdrxR16       utils.ENUMERATED
	TimeoffsetEdrxShortR16 utils.ENUMERATED
	TimeoffsetEdrxLongR16  *utils.ENUMERATED
	NumdrxCyclesrelaxedR16 *utils.ENUMERATED
	PowerboostR16          *utils.ENUMERATED
}
