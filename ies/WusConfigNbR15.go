package ies

// WUS-Config-NB-r15 ::= SEQUENCE
// Extensible
type WusConfigNbR15 struct {
	MaxdurationfactorR15   WusMaxdurationfactorNbR15
	NumposR15              WusConfigNbR15NumposR15
	NumdrxCyclesrelaxedR15 WusConfigNbR15NumdrxCyclesrelaxedR15
	TimeoffsetdrxR15       WusConfigNbR15TimeoffsetdrxR15
	TimeoffsetEdrxShortR15 WusConfigNbR15TimeoffsetEdrxShortR15
	TimeoffsetEdrxLongR15  *WusConfigNbR15TimeoffsetEdrxLongR15
}
