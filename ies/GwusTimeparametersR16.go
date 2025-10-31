package ies

// GWUS-TimeParameters-r16 ::= SEQUENCE
// Extensible
type GwusTimeparametersR16 struct {
	MaxdurationfactorR16   GwusTimeparametersR16MaxdurationfactorR16
	NumposR16              GwusTimeparametersR16NumposR16
	TimeoffsetdrxR16       GwusTimeparametersR16TimeoffsetdrxR16
	TimeoffsetEdrxShortR16 GwusTimeparametersR16TimeoffsetEdrxShortR16
	TimeoffsetEdrxLongR16  *GwusTimeparametersR16TimeoffsetEdrxLongR16
	NumdrxCyclesrelaxedR16 *GwusTimeparametersR16NumdrxCyclesrelaxedR16
	PowerboostR16          *GwusTimeparametersR16PowerboostR16
}
