package ies

// BandCombinationParameters-v10i0 ::= SEQUENCE
type BandcombinationparametersV10i0 struct {
	BandparameterlistV10i0 *[]BandparametersV10i0 `lb:1,ub:maxSimultaneousBandsR10`
}
