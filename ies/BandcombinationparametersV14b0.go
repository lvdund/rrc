package ies

// BandCombinationParameters-v14b0 ::= SEQUENCE
type BandcombinationparametersV14b0 struct {
	BandparameterlistV14b0 *[]BandparametersV14b0 `lb:1,ub:maxSimultaneousBandsR10`
}
